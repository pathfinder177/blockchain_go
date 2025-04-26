package tcp_server

import (
	"log"
	"net"
)

const (
	nodeAddress string = "127.0.0.1"
)

func Start(tcpPort string) {
	log.Printf("Server is listening on localhost%s\n", tcpPort)
	ln, err := net.Listen("tcp", nodeAddress+tcpPort)
	if err != nil {
		log.Fatal("net.Listen:", err)
	}
	defer ln.Close()

	for {
		_, err := ln.Accept()
		if err != nil {
			log.Panic(err)
		}
	}
}
