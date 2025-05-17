package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
	"wallet_server/internal/entity"
)

type Server struct {
	mux    *http.ServeMux
	server *http.Server
}

func NewServer(listenAddr string) *Server {
	m := http.NewServeMux()

	return &Server{
		mux:    m,
		server: &http.Server{Addr: listenAddr, Handler: m},
	}
}

func (router *Router) sendCurrencyHandler(w http.ResponseWriter, r *http.Request) {
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

	wallet := entity.Wallet{Address: req.Sender}
	result, err := router.UCSendCurrency.SendCurrency(
		ctx, wallet,
		req.Amount,
		req.Currency,
		req.Receiver,
		req.Mine,
	)
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

func (router *Router) gTXHistoryHandler(w http.ResponseWriter, r *http.Request) {
	type TXHistoryResponse struct {
		History string `json:"history"`
	}

	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Missing 'address' parameter", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	e := entity.Wallet{Address: address}
	history, err := router.UCGetTransactionsHistory.GetHistory(ctx, e)
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

func (router *Router) gWBHandler(w http.ResponseWriter, r *http.Request) {
	type WalletBalanceResponse struct {
		Address string `json:"address"`
		Balance string `json:"balance"`
	}

	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Missing 'address' parameter", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	e := entity.Wallet{Address: address}
	wb, err := router.UCGetBalance.GetBalance(ctx, e)
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

func (s *Server) Start(router *Router) {
	s.mux.HandleFunc("/get_wallet_balance", router.gWBHandler)
	s.mux.HandleFunc("/get_transactions_history", router.gTXHistoryHandler)

	s.mux.HandleFunc("/send_currency", router.sendCurrencyHandler)

	log.Printf("HTTPServer is listening on http://%s\n", s.server.Addr)

	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("ListenAndServe:", err)
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
