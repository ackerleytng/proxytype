package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/godbus/dbus/v5"
	"golang.org/x/sys/unix"

	evdev "github.com/gvalkov/golang-evdev"
)

func readKeys(input string) (uint8, [6]uint8) {
	inputs := strings.Split(input, " ")

	modKeys64, err := strconv.ParseInt(inputs[0], 10, 9)
	if err != nil {
		modKeys64 = 0
	}
	modKeys := uint8(modKeys64)

	var keys [6]uint8

	for i, s := range inputs[1:] {
		if i > 5 {
			break
		}

		parsed, err := strconv.ParseInt(s, 10, 9)
		if err != nil {
			log.Println("Can't parse", s)
			keys[i] = 0
		} else {
			keys[i] = uint8(parsed)
		}
	}

	return modKeys, keys
}

func buildKeysToSend(keyPressState *keyPressState) [10]uint8 {
	var out [10]uint8

	out[0] = 0xa1
	out[1] = 0x1
	out[2] = keyPressState.ModKeys

	for i, k := range keyPressState.Keys {
		out[i+4] = k
	}

	return out
}

func readEvent(dev *evdev.InputDevice) (*evdev.InputEvent, bool) {
	if err := dev.File.SetReadDeadline(time.Now().Add(250 * time.Millisecond)); err != nil {
		log.Fatalln("Can't SetReadDeadline")
	}

	event, err := dev.ReadOne()
	if err != nil {
		if strings.Contains(err.Error(), "i/o timeout") {
			return event, false
		} else {
			log.Fatalln("Other error reading", err)
		}
	}

	return event, true

}

type keyPressState struct {
	// A single uint8 holding modifier keys
	ModKeys uint8
	// An array of 6 uint8s holding all the keys that are pressed
	Keys [6]uint8
}

func addKey(keyPressState *keyPressState, scanCode uint8) {
	if mask, ok := ModKeyToMask[scanCode]; ok {
		keyPressState.ModKeys |= mask
	} else {
		if keyCode, ok := ScanCodeToKeyCode[scanCode]; ok {
			addToOrderedSet(&keyPressState.Keys, keyCode)
		} else {
			log.Println("Add - unknown scancode", scanCode)
		}
	}
}

func removeKey(keyPressState *keyPressState, scanCode uint8) {
	if mask, ok := ModKeyToMask[scanCode]; ok {
		keyPressState.ModKeys &= ^mask
	} else {
		if keyCode, ok := ScanCodeToKeyCode[scanCode]; ok {
			removeFromOrderedSet(&keyPressState.Keys, keyCode)
		} else {
			log.Println("Remove - unknown scancode", scanCode)
		}
	}
}

func handleKeyEvent(socket int, keyPressState *keyPressState, event *evdev.InputEvent) {
	switch keyEvent := evdev.NewKeyEvent(event); keyEvent.State {
	case evdev.KeyDown:
		addKey(keyPressState, uint8(keyEvent.Scancode))
	case evdev.KeyUp:
		removeKey(keyPressState, uint8(keyEvent.Scancode))
	default:
		log.Println("Other evdev event", keyEvent.State)
		return
	}

	keysToSend := buildKeysToSend(keyPressState)
	log.Println("Sending", keysToSend)
	unix.Send(socket, keysToSend[:], 0)
}

func handleKeyEvents(socket int, dev *evdev.InputDevice) {
	if err := dev.Grab(); err != nil {
		log.Fatalln("Can't grab device")
	}
	defer dev.Release()

	syscall.SetNonblock(int(dev.File.Fd()), true)

	dev.SetRepeatRate(200, 200)

	log.Println("Start using", dev.Name)
	var keyPressState keyPressState
	for {
		event, ok := readEvent(dev)
		if ok && event.Type == evdev.EV_KEY {
			handleKeyEvent(socket, &keyPressState, event)
		}
	}
}

func getUserDeviceChoice() *evdev.InputDevice {
	devices, err := evdev.ListInputDevices()
	if err != nil {
		log.Fatalln("Can't list devices")
	}

	for i, d := range devices {
		fmt.Printf("%3v: %v\n", i, d.Name)
	}

	var inputString string
	for {
		fmt.Println("Pick keyboard to proxy to bluetooth:")
		fmt.Scanln(&inputString)

		input, err := strconv.ParseInt(inputString, 10, 32)
		if err != nil {
			fmt.Println("Error parsing", inputString)
			fmt.Println("Try again, enter a number")
			continue
		}

		if int(input) < len(devices) {
			return devices[input]
		} else {
			fmt.Println("Pick a number less than", len(devices))
		}
	}
}

func startServer(address [6]byte) {
	controlSocket := bindSocket(address, 0x11)
	defer closeSocket(controlSocket)
	interruptSocket := bindSocket(address, 0x13)
	defer closeSocket(interruptSocket)

	log.Println("Bound both ports")

	acceptedControlSocket := acceptAndPrint(controlSocket, 0x11)
	defer closeSocket(acceptedControlSocket)
	acceptedInterruptSocket := acceptAndPrint(interruptSocket, 0x13)
	defer closeSocket(acceptedInterruptSocket)

	dev := getUserDeviceChoice()
	handleKeyEvents(acceptedInterruptSocket, dev)
}

func handleExit(conn *dbus.Conn) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	go func() {
		select {
		case <-c:
			unexportObject(conn)
			conn.Close()

			os.Exit(0)
		}
	}()
}

func getBluetoothInterfaceAddress(conn *dbus.Conn) [6]byte {
	var address string

	call := conn.Object("org.bluez", "/org/bluez/hci0").Call("org.freedesktop.DBus.Properties.Get", 0, "org.bluez.Adapter1", "Address")
	if err := call.Store(&address); err != nil {
		log.Fatal("Can't get bluetooth interface address", err, address)
	}

	addrRaw, err := hex.DecodeString(strings.ReplaceAll(address, ":", ""))
	if err != nil {
		log.Fatalln("Can't decode string", address)
	}

	var addr [6]byte
	copy(addr[:], addrRaw)
	return addr
}

func main() {
	conn := connectDbus()
	defer conn.Close()
	handleExit(conn)

	// If an application disconnects from the bus, all its registered profiles will be removed
	registerProfile(conn)

	address := getBluetoothInterfaceAddress(conn)
	startServer(address)
}
