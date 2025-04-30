package app

import (
	"context"
	"fmt"
	"time"
	"wallet_server/internal/entity"
	"wallet_server/internal/gateway/tcp"
	GetTransactionsHistoryInteractor "wallet_server/internal/usecase/getTransactionsHistoryInteractor"
)

const (
	appServerAddr string = "localhost"
	appPort       string = ":3003"

	tcpServerAddr string = "localhost"
	tcpServerPort string = ":4000"

	walletNodeAddress string = "localhost"
	walletNodePort    string = ":3001"
)

func Run() {
	e := entity.Wallet{Address: "1DQz9eQSGRuqfk7npzDzdz53DSKEbHEUGV"}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	gateway := tcp.New(tcpServerAddr+tcpServerPort, walletNodeAddress+walletNodePort)
	ucGTHI := GetTransactionsHistoryInteractor.New(gateway)

	h, _ := ucGTHI.GetHistory(ctx, e)
	fmt.Println(h)

	//gateway
	// go tcp.Listen(tcpServerAddr, tcpPort)

	//usecases

	//controller
	// go http_server.Start(appServerAddr, appPort)

	// for {
	// 	time.Sleep(time.Second * 300)
	// } //FIXME

	//graceful shutdown here
}
