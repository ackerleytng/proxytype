package main

import (
	"io/ioutil"
	"log"

	"github.com/godbus/dbus/v5"
)

const profilePath = dbus.ObjectPath("/bluez/clipcontrol")

// Add 0x1124 to BASE_UUID of 00000000-0000-1000-8000-00805f9b34fb
// https://btprodspecificationrefs.blob.core.windows.net/assigned-numbers/Assigned%20Number%20Types/Service%20Discovery.pdf
// "00001124-0000-1000-8000-00805f9b34fb"
const hidServiceUuid = "1c2680e3-17b2-457a-a981-d680d2387255"

type clipControlExport struct{}

func (export clipControlExport) Release() *dbus.Error {
	log.Println("Release handler called")
	return nil
}

func (export clipControlExport) NewConnection(path dbus.ObjectPath, fd dbus.UnixFD, fdProperties map[string]dbus.Variant) *dbus.Error {
	log.Println("NewConnection handler called")
	return nil
}

func (export clipControlExport) RequestDisconnection(path dbus.ObjectPath) *dbus.Error {
	log.Println("RequestDisconnection handler called")
	return nil
}

func readSdpRecord() dbus.Variant {
	data, err := ioutil.ReadFile("./sdp_record.xml")
	if err != nil {
		log.Fatal(err)
	}

	return dbus.MakeVariant(string(data))
}

func unregisterProfile(conn *dbus.Conn) string {
	var s string
	call := conn.Object("org.bluez", "/org/bluez").Call("org.bluez.ProfileManager1.UnregisterProfile", 0, profilePath)
	if err := call.Store(&s); err.Error() != "dbus.Store: length mismatch" {
		log.Fatal("UnregisterProfile:", err)
	}

	return s
}

func unexportObject(conn *dbus.Conn) {
	conn.Export(nil, profilePath, "org.bluez.Profile1")
	log.Println("Unexported object")
}

func registerProfile(conn *dbus.Conn) {
	opts := map[string]dbus.Variant{
		"Role":                  dbus.MakeVariant("server"),
		"RequireAuthentication": dbus.MakeVariant(true),
		"RequireAuthorization":  dbus.MakeVariant(true),
		"AutoConnect":           dbus.MakeVariant(true),
		"ServiceRecord":         readSdpRecord(),
		"PSM":                   dbus.MakeVariant(0x11),
	}

	// A call to RegisterProfile doesn't return anything
	conn.Object("org.bluez", "/org/bluez").Call("org.bluez.ProfileManager1.RegisterProfile", 0, profilePath, hidServiceUuid, opts)
	log.Println("Registered profile")

	conn.Export(&clipControlExport{}, profilePath, "org.bluez.Profile1")
	log.Println("Exported object")
}

func connectDbus() *dbus.Conn {
	conn, err := dbus.ConnectSystemBus()
	if err != nil {
		log.Fatalln("Failed to connect to SystemBus bus:", err)
	}

	return conn
}
