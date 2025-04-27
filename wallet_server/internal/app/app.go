package app

import (
	"time"
	"wallet_server/internal/controller/http_server"
	"wallet_server/internal/controller/tcp_server"
)

const (
	appServerAddr string = "http://localhost"
	appPort       string = ":3003"

	tcpServerAddr string = "http://localhost"
	tcpPort       string = ":4000"

	walletNodeAddress string = "http://localhost"
	walletNodePort    string = ":3001"
)

func Run() {
	//gateway

	//usecases

	//controller
	go http_server.Start(appServerAddr, appPort)
	go tcp_server.Start(tcpServerAddr, tcpPort)

	for {
		time.Sleep(time.Second * 300)
	} //FIXME

	//graceful shutdown here
}
