package app

import (
	"time"
	"wallet_server/internal/controller/http_server"
	"wallet_server/internal/controller/tcp_server"
)

const (
	appPort string = ":3003"
	tcpPort string = ":4000"

	walletNodeAddress string = "http://localhost"
	walletNodePort    string = ":3001"
)

func Run() {
	//repo

	//usecases

	//controller
	go http_server.Start(appPort)
	go tcp_server.Start(tcpPort)

	for {
		time.Sleep(time.Second * 300)
	} //FIXME

	//graceful shutdown here
}
