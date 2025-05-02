package app

import (
	"time"
	"wallet_server/internal/controller/server"
)

const (
	appServerAddr string = "localhost:3003"

	blockchainNodeAddress string = "localhost:3000"

	tcpServerAddr string = "localhost:4000"
)

func Run() {
	//usecases

	//controller
	go server.Start(appServerAddr)

	for {
		time.Sleep(time.Second * 300)
	} //FIXME

	//graceful shutdown here
}
