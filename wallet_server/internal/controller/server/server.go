package server

import (
	"encoding/json"
	"log"
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
	//interactor -> cliGateway
	wb, err := getWalletBalance(address)
	if wb == "" {
		http.Error(w, "Wallet balance is empty", http.StatusInternalServerError)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//interactor returns
	response := WalletBalanceResponse{
		Address: address,
		Balance: wb,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getWalletBalance(address string) (string, error) {
	args := []string{"getbalance", "-address", address}
	cmd := exec.Command("./blockchain", args...)
	cmd.Env = append(os.Environ(), "NODE_ID=3001")
	cmd.Dir = "/home/pathfinder177/projects/blockchain/cmd/app"

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func StartServer(appPort string) {
	http.HandleFunc("/get_wallet_balance", gWBHandler)

	log.Printf("Server is listening on http://localhost%s\n", appPort)
	if err := http.ListenAndServe(appPort, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
