package main

import (
	"log"

	"golang.org/x/sys/unix"
)

func bindSocket(address [6]byte, port uint16) int {
	// SOCK_SEQPACKET and SOCK_STREAM are both able to allow connections fine
	s, err := unix.Socket(unix.AF_BLUETOOTH, unix.SOCK_SEQPACKET, unix.BTPROTO_L2CAP)
	if err != nil {
		log.Fatalln("Failed to create socket for port", port, "err=", err)
	}

	// If SO_REUSEADDR is not set, this will still work?
	if err := unix.SetsockoptInt(s, unix.SOL_SOCKET, unix.SO_REUSEADDR, 1); err != nil {
		log.Fatalln("Failed to setsockopt SO_REUSEADDR on port", port, "err=", err)
	}

	// If Addr is not specified, the default bluetooth interface will be used?
	if err := unix.Bind(s, &unix.SockaddrL2{PSM: port, Addr: address}); err != nil {
		log.Fatalln("Failed to bind socket on port", port, "err=", err)
	}

	if err := unix.Listen(s, 1); err != nil {
		log.Fatalln("Failed to listen on port", port, "err=", err)
	}

	return s
}

func acceptAndPrint(socket int, port int) int {
	acceptedSocket, addr, err := unix.Accept(socket)
	if err != nil {
		log.Fatalln("Failed to accept")
	}
	log.Println("Accepted on port", port, "from addr", addr)

	return acceptedSocket
}

func closeSocket(socket int) {
	unix.Close(socket)
}
