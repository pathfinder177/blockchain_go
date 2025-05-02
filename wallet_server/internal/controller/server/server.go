package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
	"wallet_server/internal/entity"
	"wallet_server/internal/gateway/cli"
	"wallet_server/internal/gateway/tcp"
	GetBalanceInteractor "wallet_server/internal/usecase/getBalanceInteractor"
	GetTransactionsHistoryInteractor "wallet_server/internal/usecase/getTransactionsHistoryInteractor"
	SendCurrencyInteractor "wallet_server/internal/usecase/sendCurrencyInteractor"
)

func sendCurrencyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Amount   string `json:"amount"`
		Currency string `json:"currency"`
		Sender   string `json:"sender"`
		Receiver string `json:"receiver"`
		Mine     string `json:"mine"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid JSON"}`, http.StatusBadRequest)
		return
	}

	if req.Amount == "" || req.Currency == "" ||
		req.Sender == "" || req.Receiver == "" || req.Mine == "" {
		http.Error(w, `{"error":"missing required fields"}`, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	//FIXME to app(inject in func)
	gateway := cli.New()
	ucsci := SendCurrencyInteractor.New(gateway)
	//

	wallet := entity.Wallet{Address: req.Sender}
	result, err := ucsci.SendCurrency(ctx, wallet, req.Amount, req.Currency, req.Receiver, req.Mine)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}
	splitted_res := strings.Split(result, "\n")

	w.Header().Set("Content-Type", "application/json")
	response := struct {
		SendResult string `json:"sendResult"`
	}{SendResult: splitted_res[len(splitted_res)-1]}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		// If even this fails, we’re in a bad state—log the error
		log.Printf("encode response: %v", err)
	}

}

func gTXHistoryHandler(w http.ResponseWriter, r *http.Request) {
	type TXHistoryResponse struct {
		History string `json:"history"`
	}

	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Missing 'address' parameter", http.StatusBadRequest)
		return
	}
	e := entity.Wallet{Address: address}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	//fixme to app(inject into func)
	tcpServerAddr := "localhost:4000"
	blockchainNodeAddr := "localhost:3000"

	gateway := tcp.New(tcpServerAddr, blockchainNodeAddr)
	ucGTHI := GetTransactionsHistoryInteractor.New(gateway)
	//

	history, err := ucGTHI.GetHistory(ctx, e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := TXHistoryResponse{
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

func Start(appServerAddr string) {
	http.HandleFunc("/get_wallet_balance", gWBHandler)
	http.HandleFunc("/get_transactions_history", gTXHistoryHandler)
	http.HandleFunc("/send_currency", sendCurrencyHandler)

	log.Printf("HTTPServer is listening on http://%s\n", appServerAddr)
	if err := http.ListenAndServe(appServerAddr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
