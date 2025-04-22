package main

import (
	"encoding/json"
	"net/http"
	"os"
	"os/exec"
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

	wb, err := getWalletBalance(address)
	if wb == "" {
		http.Error(w, "Wallet balance is empty", http.StatusInternalServerError)
		return
	} else if err != nil {
		http.Error(w, "Wallet address is not correct", http.StatusBadRequest)
		return
	}

	response := WalletBalanceResponse{
		Address: address,
		Balance: wb,
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Encode response to JSON and write it
	json.NewEncoder(w).Encode(response)
}

func getWalletBalance(address string) (string, error) {
	args := []string{"getbalance", "-address", address}
	cmd := exec.Command("./blockchain", args...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func main() {
	os.Setenv("NODE_ID", "3001") //FIXME

	http.HandleFunc("/get_wallet_balance", gWBHandler)
	http.ListenAndServe(":3003", nil)
}

// package main

// import "blockchain/internal/app"

// func main() {
// 	cli := app.CLI{}
// 	cli.Run()
// }
