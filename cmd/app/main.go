package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"text/template"
)

var index_tmpl_get = template.Must(template.New("form").Parse(`
<!DOCTYPE html>
<html>
<head>
    <title>Input Form</title>
</head>
<body>
    <h1>Enter your wallet address</h1>
    <form action="/" method="POST">
        <input type="text" name="walletAddress" required>
        <input type="submit" value="Submit">
    </form>
</body>
</html>
`))

var index_tmpl_post = template.Must(template.New("form").Parse(`
<!DOCTYPE html>
<html>
<head><title>Welcome</title></head>
<body>
    <h1>Your wallet address is {{.WAddress}}</h1>
	<hr>

    <h2>Balance</h2>
    {{range .WBalance}}
        <p>{{.}}</p>
    {{end}}
	<hr>

    <h2>Actions</h2>
	<form action="/transactions" method="GET">
        <button type="submit">Get Transactions History</button><br />
    </form>
    <form>
        <button type="button">Send Currency</button><br />
        <button type="button">Get Currency Transactions History</button><br />
        <button type="button">Delete Wallet</button><br />
    </form>
</body>
</html>
`))

func index(w http.ResponseWriter, r *http.Request) { //DELETEME and templates
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}
		address := r.FormValue("walletAddress")
		wb, err := getWalletBalance(address)
		if err != nil {
			http.Error(w, "Wallet address is not correct", http.StatusBadRequest)
			return
		}

		index_tmpl_post.Execute(w, struct {
			WAddress string
			WBalance []string
		}{address, wb})
		return
	}

	index_tmpl_get.Execute(w, nil)
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

	wb, err := getWalletBalance(address)
	if err != nil {
		http.Error(w, "Wallet address is not correct", http.StatusBadRequest)
		return
	}
	b := strings.Join(wb, "") // or use a separator like " " if needed

	response := WalletBalanceResponse{
		Address: address,
		Balance: b,
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Encode response to JSON and write it
	json.NewEncoder(w).Encode(response)
}

func getWalletBalance(address string) ([]string, error) {
	args := []string{"getbalance", "-address", address}
	cmd := exec.Command("./blockchain", args...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return []string{}, err
	}

	strSliceOutput := []string{}
	lastPos := 0
	for i := range output {
		if output[i] == '\n' {
			s := string(output[lastPos : i+1])
			strSliceOutput = append(strSliceOutput, s)
			lastPos = i + 1
		}
	}

	return strSliceOutput, nil
}

func transactionsHistory(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", "transactions history!")
	//getTH
}

func main() {
	// os.Setenv("NODE_ID", "3003") //FIXME use it only on server side
	// startWalletServer()
	// getTransactionsHistory("14tmM4cbsoMqJvMv2dixauXFxKRaKnibad")

	http.HandleFunc("/get_wallet_balance", gWBHandler)

	http.ListenAndServe(":3003", nil)
}

// package main

// import "blockchain/internal/app"

// func main() {
// 	cli := app.CLI{}
// 	cli.Run()
// }
