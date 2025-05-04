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
	appServerAddr      string = "localhost:3003"
	blockchainNodeAddr string = "localhost:3000"
	tcpServerAddr      string = "localhost:4000"
)

func Run() {
	//gateway
	cliGateway := cli.New()
	tcpGateway := tcp.New(tcpServerAddr, blockchainNodeAddr)

	//usecases
	UCGetBalance := GetBalanceInteractor.New(cliGateway)
	UCGetTransactionsHistory := GetTransactionsHistoryInteractor.New(tcpGateway)

	UCSendCurrency := SendCurrencyInteractor.New(cliGateway)

	//controller
	router := server.New(
		UCGetBalance,
		UCGetTransactionsHistory,
		UCSendCurrency,
	)
	server.Start(appServerAddr, router) //FIXME go

	//graceful shutdown here
}
