package http_server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"wallet_server/internal/entity"
	"wallet_server/internal/gateway/cli"
	GetBalanceInteractor "wallet_server/internal/usecase/getBalanceInteractor"
)

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

	log.Printf("HTTPServer is listening on http://localhost%s\n", appPort)
	if err := http.ListenAndServe(appServerAddr+appPort, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
