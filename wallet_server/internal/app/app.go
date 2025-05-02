package app

import (
	"time"
	"wallet_server/internal/controller/http_server"
)

const (
	appServerAddr string = "localhost"
	appPort       string = ":3003"

	blockchainNodeAddress string = "localhost"
	blockchainNodePort    string = ":3000"

	tcpServerAddr string = "localhost"
	tcpServerPort string = ":4000"
)

func Run() {
	//usecases

	//controller
	go http_server.Start(appServerAddr, appPort)

	for {
		time.Sleep(time.Second * 300)
	} //FIXME

	//graceful shutdown here
}
