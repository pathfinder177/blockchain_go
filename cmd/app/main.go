package main

import (
	"log"
	"net/http"
	"os/exec"
	"text/template"
)

var tmpl = template.Must(template.New("form").Parse(`
<!DOCTYPE html>
<html>
<head><title>Input Form</title></head>
<body>
    <h1>Enter your wallet address</h1>
    <form action="/" method="POST">
        <input type="text" name="walletAddress" required>
        <input type="submit" value="Submit">
    </form>
    {{if .}}
        <p>Your wallet address and balance is {{.}}</p> //FIXME
    {{end}}
</body>
</html>
`))

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}
		address := r.FormValue("walletAddress")
		wb, err := getWalletBalance(address)
		if err != nil {
			log.Panic(err)
		}
		tmpl.Execute(w, address+wb) //FIXME
		return
	}

	tmpl.Execute(w, nil)
}

func getWalletBalance(address string) (string, error) {
	args := []string{"getbalance", "-address", address}
	cmd := exec.Command("./blockchain", args...)

	output, err := cmd.CombinedOutput() // gets both stdout and stderr
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func main() {
	http.HandleFunc("/", index)

	http.ListenAndServe(":3003", nil)
}

// package main

// import "blockchain/internal/app"

// func main() {
// 	cli := app.CLI{}
// 	cli.Run()
// }
