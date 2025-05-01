package http_server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"wallet_server/internal/entity"
	"wallet_server/internal/gateway/cli"
	"wallet_server/internal/gateway/tcp"
	GetBalanceInteractor "wallet_server/internal/usecase/getBalanceInteractor"
	GetTransactionsHistoryInteractor "wallet_server/internal/usecase/getTransactionsHistoryInteractor"
)

func gTXHistoryHandler(w http.ResponseWriter, r *http.Request) {
	type TXHistoryResponse struct {
		Address string `json:"address"`
		History string `json:"history"`
	}

	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Missing 'address' parameter", http.StatusBadRequest)
		return
	}
	e := entity.Wallet{Address: address}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	//fixme to app(inject into func)
	tcpServerAddr := "localhost"
	tcpServerPort := ":4000"
	walletNodeAddress := "localhost"
	walletNodePort := ":3001"

	gateway := tcp.New(tcpServerAddr+tcpServerPort, walletNodeAddress+walletNodePort)
	ucGTHI := GetTransactionsHistoryInteractor.New(gateway)
	//

	history, err := ucGTHI.GetHistory(ctx, e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := TXHistoryResponse{
		Address: address,
		History: history,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func gWBHandler(w http.ResponseWriter, r *http.Request) {
	type WalletBalanceResponse struct {
		Address string `json:"address"`
		Balance string `json:"balance"`
	}

	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Missing 'address' parameter", http.StatusBadRequest)
		return
	}
	e := entity.Wallet{Address: address}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	//FIXME to app(inject in func)
	gateway := cli.New()
	ucgbi := GetBalanceInteractor.New(gateway)
	//

	wb, err := ucgbi.GetBalance(ctx, e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := WalletBalanceResponse{
		Address: address,
		Balance: wb,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func Start(appServerAddr, appPort string) {
	http.HandleFunc("/get_wallet_balance", gWBHandler)
	http.HandleFunc("/get_transactions_history", gTXHistoryHandler)

	log.Printf("HTTPServer is listening on http://localhost%s\n", appPort)
	if err := http.ListenAndServe(appServerAddr+appPort, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
