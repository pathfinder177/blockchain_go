package app

import (
	"time"
	"wallet_server/internal/controller/http_server"
	"wallet_server/internal/gateway/tcp"
)

const (
	appServerAddr string = "localhost"
	appPort       string = ":3003"

	tcpServerAddr string = "localhost"
	tcpPort       string = ":4000"

	walletNodeScheme  string = "http://"
	walletNodeAddress string = "localhost"
	walletNodePort    string = ":3001"
)

func Run() {
	//gateway
	go tcp.Listen(tcpServerAddr, tcpPort)

	//usecases

	//controller
	go http_server.Start(appServerAddr, appPort)

	for {
		time.Sleep(time.Second * 300)
	} //FIXME

	//graceful shutdown here
}
