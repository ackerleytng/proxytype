# proxytype

Convert your computer (raspberry pi, etc) into a bluetooth keyboard!

proxytype is a service that reads all keystrokes on an input device and proxies it to another device through bluetooth.

## Usage

This has only been tested with Linux, and depends on dbus and the bluez stack

### 1. Start `bluetoothd` without the input plugin

The input plugin listens on the same ports that proxytype listens to, hence we have to disable that by starting `bluetoothd` without the input plugin with

```
sudo /usr/lib/bluetooth/bluetoothd --noplugin input
```

### 2. Put the bluetooth adapter in listening mode

```
$ bluetoothctl
[bluetooth]# discoverable on
```

### 3. Compile and start proxytype

A successful run should look something like this

```
$ go build && sudo ./proxytype
[sudo] password for ackerleytng:
2021/10/03 21:45:36 Registered profile
2021/10/03 21:45:36 Exported object
2021/10/03 21:45:36 Bound both ports
2021/10/03 21:45:52 Accepted on port 17 from addr &{17 65 [78 166 78 110 183 96] 0 {0 0 [0 0 0 0 0 0] 0 0 [0]}}
2021/10/03 21:45:52 Accepted on port 19 from addr &{19 66 [78 166 78 110 183 96] 0 {0 0 [0 0 0 0 0 0] 0 0 [0]}}
  0: Sleep Button
  1: Lid Switch
  2: HDA Intel PCH Mic
  3: HDA Intel PCH Headphone
  4: HDA Intel PCH HDMI/DP,pcm=3
  5: HDA Intel PCH HDMI/DP,pcm=7
  6: HDA Intel PCH HDMI/DP,pcm=8
  7: HDA Intel PCH HDMI/DP,pcm=9
  8: HDA Intel PCH HDMI/DP,pcm=10
  9: CX 2.4G Receiver
 10: CX 2.4G Receiver Mouse
 11: CX 2.4G Receiver
 12: Power Button
 13: CX 2.4G Receiver Consumer Control
 14: CX 2.4G Receiver System Control
 15: AT Translated Set 2 keyboard
 16: ThinkPad Extra Buttons
 17: PC Speaker
 18: Video Bus
 19: Synaptics TM3289-021
 20: Integrated Camera: Integrated C
 21: TPPS/2 Elan TrackPoint
Pick keyboard to proxy to bluetooth:
9
2021/10/03 21:45:59 Start using CX 2.4G Receiver
2021/10/03 21:46:12 Sending [161 1 0 0 23 0 0 0 0 0]
2021/10/03 21:46:12 Sending [161 1 0 0 0 0 0 0 0 0]
```

After picking the input device of your choice, the input device will be interpreted as a keyboard, and all keystrokes will be sent over bluetooth to the connected bluetooth device.

I didn't build any functionality for exiting, so the proxied device can't be used to exit the program.

I usually use another keyboard (not the one being proxied) to send `SIGINT` to proxytype.

## Overview

Here's what proxytype does

+ Advertise a keyboard and how to connect to this "keyboard" over bluetooth. The properties of this "keyboard" are defined in `sdp_record.xml`
    + This advertisement is done by registering a profile over dbus (`org.bluez.ProfileManager1.RegisterProfile`)
    + The advertisement is not complete without also exporting a profile object of type `org.bluez.Profile1`
+ Having a server listen for connections on the ports as specified in the properties in the advertisement
    + After accepting a connection, ignore all inputs from the connection
    + Send keycodes over the interrupt channel
+ Reads inputs using evdev, translates those scancodes to keycodes, sends the keycodes over the interrupt channel

## Acknowledgements

+ [Bluetooth Essentials for Programmers by Albert S. Huang and Larry Rudolph](https://doi.org/10.1017/CBO9780511546976)
    + This book has to be the best overview of bluetooth programming ever! This helped me understand bluez the various interactions between client and server.
+ [Raspberry Pi Bluetooth HID gist](https://gist.github.com/ukBaz/a47e71e7b87fbc851b27cde7d1c0fcf0)
    + This gist and the super helpful and friendly community around the gist
+ [go-hidproxy](https://github.com/rosmo/go-hidproxy/)
    + This repo provided another sample, in golang
+ Gonna do a shoutout for [zephyr](https://docs.zephyrproject.org/latest/guides/bluetooth/bluetooth-tools.html#running-on-qemu-and-native-posix) too!
    + If I were to do this again from the beginning I would try zephyr first; it seems to be a more direct approach than bluez+dbus
