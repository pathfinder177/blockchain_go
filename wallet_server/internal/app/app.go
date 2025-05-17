package app

import (
	"wallet_server/internal/controller/server"
	"wallet_server/internal/gateway/cli"
	"wallet_server/internal/gateway/tcp"
	GetBalanceInteractor "wallet_server/internal/usecase/getBalanceInteractor"
	GetTransactionsHistoryInteractor "wallet_server/internal/usecase/getTransactionsHistoryInteractor"
	SendCurrencyInteractor "wallet_server/internal/usecase/sendCurrencyInteractor"
)

const (
	blockchainNodeAddr string = "localhost:3000"
	listenAddr         string = "localhost:3003"
	tcpServerAddr      string = "localhost:4000"
)

func Run() {
	//FIXME add config

	//gateway
	cliGateway := cli.New()
	tcpGateway := tcp.New(tcpServerAddr, blockchainNodeAddr)

	//usecases
	UCGetBalance := GetBalanceInteractor.New(cliGateway)
	UCGetTransactionsHistory := GetTransactionsHistoryInteractor.New(tcpGateway)

	UCSendCurrency := SendCurrencyInteractor.New(cliGateway)

	//controller
	router := server.NewRouter(
		UCGetBalance,
		UCGetTransactionsHistory,
		UCSendCurrency,
	)
	server := server.NewServer(listenAddr)

	server.Start(router) //FIXME go

	//graceful shutdown
}
