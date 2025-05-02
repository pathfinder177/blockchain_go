package tcp

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"time"
	"wallet_server/internal/entity"
)

var blocksInTransit = [][]byte{}

type block struct {
	AddrFrom string
	Block    []byte
}

type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
	Height        int
}

type Transaction struct {
	ID       []byte
	Vin      []TXInput
	Vout     []TXOutput
	Currency string
}

type TXInput struct {
	Txid      []byte
	Vout      int
	Signature []byte
	PubKey    []byte
}

type TXOutput struct {
	Value      int
	PubKeyHash []byte
}

type getblocks struct {
	AddrFrom string
}

type getdata struct {
	AddrFrom string
	Type     string
	ID       []byte
}

type inv struct {
	AddrFrom string
	Type     string
	Items    [][]byte
}

func (tcpGateway *tcpGateway) sendData(addr string, data []byte) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Printf("%s is not available\n", addr)
		return
	}
	defer conn.Close()

	_, err = io.Copy(conn, bytes.NewReader(data))
	if err != nil {
		log.Panic(err)
	}
}

func (tcpGateway *tcpGateway) sendGetData(address, kind string, id []byte) {
	payload := gobEncode(getdata{tcpGateway.thisAddr, kind, id})
	request := append(commandToBytes("getdata"), payload...)

	tcpGateway.sendData(address, request)
}

func (tcpGateway *tcpGateway) handleInv(request []byte) {
	var buff bytes.Buffer
	var payload inv

	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Recevied inventory with %d %s\n", len(payload.Items), payload.Type)

	if payload.Type == "block" {
		blocksInTransit = payload.Items

		blockHash := payload.Items[0]
		tcpGateway.sendGetData(payload.AddrFrom, "block", blockHash)

		newInTransit := [][]byte{}
		for _, b := range blocksInTransit {
			if !bytes.Equal(b, blockHash) {
				newInTransit = append(newInTransit, b)
			}
		}
		blocksInTransit = newInTransit
	}
}

func (tcpGateway *tcpGateway) handleBlock(request []byte, input chan<- *Block) {
	var buff bytes.Buffer
	var payload block

	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	blockData := payload.Block
	input <- DeserializeBlock(blockData)

	fmt.Println("Recevied a new block!") //FIXME

	if len(blocksInTransit) > 0 {
		blockHash := blocksInTransit[0]
		tcpGateway.sendGetData(payload.AddrFrom, "block", blockHash)

		blocksInTransit = blocksInTransit[1:]
	} else {
		close(input)
	}
}

func (tcpGateway *tcpGateway) getWalletTxFromBlock(tx *Transaction, timestamp int64, WAddress, WPubKeyHash string) *entity.HistoricalTransaction {
	htx := &entity.HistoricalTransaction{}

	//Sent TX FIXME
	vin := tx.Vin[0]

	sVinPK := string(vin.PubKey)
	if sVinPK == WPubKeyHash {
		htx.From = WAddress
		htx.Currency = tx.Currency
		htx.Timestamp = timestamp

		for _, vout := range tx.Vout {
			sVoutPK := string(vout.PubKeyHash)
			if sVoutPK != WPubKeyHash {
				htx.To = getWalletAddrByPubKeyHash(sVoutPK)
				htx.Amount = vout.Value

				return htx
			}
		}

	}
	//Received TX
	for _, vout := range tx.Vout {
		sVoutPK := string(vout.PubKeyHash)
		if sVoutPK == WPubKeyHash {
			htx.From = getWalletAddrByPubKeyHash(sVinPK)
			htx.To = WAddress
			htx.Currency = tx.Currency
			htx.Timestamp = timestamp
			htx.Amount = vout.Value

			return htx
		}
	}

	return nil
}

func (tcpGateway *tcpGateway) getBlocks() {
	payload := gobEncode(getblocks{tcpGateway.thisAddr})
	request := append(commandToBytes("getblocks"), payload...)

	tcpGateway.sendData(tcpGateway.otherAddr, request)
}

func (tcpGateway *tcpGateway) handleConnection(conn net.Conn, input chan<- *Block) {
	defer conn.Close()

	request, err := io.ReadAll(conn)
	if err != nil {
		log.Panic(err)
	}
	command := bytesToCommand(request[:commandLength])
	fmt.Printf("Received %s command\n", command)

	switch command {
	case "block":
		tcpGateway.handleBlock(request, input)
	case "inv":
		tcpGateway.handleInv(request)
	}
}

func (tcpGateway *tcpGateway) _getHistory(WAddress, WPubKeyHash string) ([]*entity.HistoricalTransaction, error) {
	ln, err := net.Listen("tcp", tcpGateway.thisAddr)
	if err != nil {
		log.Panic(err)
	}
	defer ln.Close()
	tcpLn := ln.(*net.TCPListener)

	input := make(chan *Block)
	output := make(chan *entity.HistoricalTransaction)
	defer close(output)

	//input reader output writer
	go func() {
		for b := range input {
			for _, tx := range b.Transactions {
				wtfb := tcpGateway.getWalletTxFromBlock(tx, b.Timestamp, WAddress, WPubKeyHash)
				if wtfb != nil {
					output <- wtfb
				}
			}
		}
	}()

	//output reader
	history := []*entity.HistoricalTransaction{}
	go func() {
		for wtfb := range output {
			history = append(history, wtfb)
		}
	}()

	tcpGateway.getBlocks() //FIXME go
	for {                  //FIXME go
		deadline := time.Now().Add(2 * time.Second)
		if err := tcpLn.SetDeadline(deadline); err != nil {
			log.Fatalf("failed to set deadline: %v", err)
		}

		conn, err := tcpLn.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Timeout() {
				fmt.Println("no connection in 2s — stopping listener")
				break
			}
			log.Panicf("accept error: %v", err)
		}

		tcpGateway.handleConnection(conn, input) //FIXME go
	}

	return history, nil
}

func (tcpGateway *tcpGateway) GetHistory(ctx context.Context, e entity.Wallet) (string, error) {
	WAddress := e.Address
	WalletPKH, err := getWalletPubKeyHash(WAddress)
	if err != nil {
		return "", err
	}
	if WalletPKH == "" {
		return "", errors.New("wallet pubkeyhash is empty")
	}

	history, err := tcpGateway._getHistory(WAddress, WalletPKH)
	if err != nil {
		return "", nil
	}

	return mapHistoryToString(history), nil
}
