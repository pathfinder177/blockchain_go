package tcp_server

import (
	"fmt"
	"io"
	"log"
	"net"
)

func GetTransactions(WAddress string) (string, error) {
	TXs := ""
	return TXs, nil
}

func handleConnection(conn net.Conn) {
	request, err := io.ReadAll(conn)
	if err != nil {
		log.Panic(err)
	}
	command := bytesToCommand(request[:commandLength])
	fmt.Printf("Received %s command\n", command)

	// switch command {
	// case "block":
	// 	handleBlock(request, chFilledBlockTxsHistory, chDone)
	// case "inv":
	// 	handleInv(request)
	// }
}

func Start(tcpServerAddr, tcpPort string) {
	log.Printf("TCPServer is listening on localhost%s\n", tcpPort)
	ln, err := net.Listen("tcp", tcpServerAddr+tcpPort)
	if err != nil {
		log.Fatal("net.Listen:", err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Panic(err)
		}

		handleConnection(conn)
	}
}
