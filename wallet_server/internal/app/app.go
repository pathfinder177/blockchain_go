package app

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"
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

	shutdownTimeout time.Duration = 5 * time.Second
)

// FIXME add config
func Run() {
	//Signal handler firstly
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

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

	go func() {
		server.Start(router)
	}()

	//graceful shutdown
	<-ctx.Done()
	log.Println("shutting down server gracefully")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	select {
	case <-shutdownCtx.Done():
		log.Fatal("shutdownTimeout")
	default:
		break
	}
}
