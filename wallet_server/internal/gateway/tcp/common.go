package tcp

import (
	"bytes"
	"cmp"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"os/exec"
	"slices"
	"strings"
	"wallet_server/internal/entity"
)

const (
	commandLength int = 12
)

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		fmt.Println("Block Deserialize error")
	}

	return &block
}

func gobEncode(data any) []byte {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(data)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// FIXME move to CLI
func getWalletAddrByPubKeyHash(WPubKeyHash string) string {
	args := []string{"getwalletaddrbypubkeyhash", "-pubkeyhash", WPubKeyHash}
	cmd := exec.Command("./blockchain", args...)
	cmd.Env = append(os.Environ(), "NODE_ID=3001")
	//FIXME
	cmd.Dir = "/home/pathfinder177/projects/blockchain/cmd/app"

	WAddress, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("can not get wallet addr by pub key hash %v", err)
	}

	return string(WAddress)
}

// FIXME move to CLI
func getWalletPubKeyHash(WAddress string) (string, error) {
	args := []string{"getwalletpubkeyhash", "-address", WAddress}
	cmd := exec.Command("./blockchain", args...)
	cmd.Env = append(os.Environ(), "NODE_ID=3001")
	//FIXME
	cmd.Dir = "/home/pathfinder177/projects/blockchain/cmd/app"

	pubKeyHash, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(pubKeyHash), nil
}

func commandToBytes(command string) []byte {
	var bytes [commandLength]byte

	for i, c := range command {
		bytes[i] = byte(c)
	}

	return bytes[:]
}

func bytesToCommand(bytes []byte) string {
	var command []byte

	for _, b := range bytes {
		if b != 0x0 {
			command = append(command, b)
		}
	}

	return string(command)
}

func sortHistoricalTX(history []*entity.HistoricalTransaction) {
	slices.SortFunc(history,
		func(a, b *entity.HistoricalTransaction) int {
			return cmp.Compare(a.Timestamp, b.Timestamp)
		})
}

func mapHistoryToString(history []*entity.HistoricalTransaction) string {
	historyToString := make([]string, len(history))
	for _, historicalTX := range history {
		s := fmt.Sprintf("%v\n", *historicalTX)
		historyToString = append(historyToString, s)
	}

	return strings.Join(historyToString, "")
}
