package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
)

// FIXME
var protocol = "tcp"
var addr = "127.0.0.1:3001"
var blocksInTransit = [][]byte{}
var nodeAddress = "127.0.0.1:3003"

const commandLength = 12

type block struct {
	AddrFrom string
	Block    []byte
}

type getblocks struct {
	AddrFrom string
}

type getdata struct {
	AddrFrom string
	Type     string
	ID       []byte
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

type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
	Height        int
}

type inv struct {
	AddrFrom string
	Type     string
	Items    [][]byte
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

// end FIXME

func getBlocks() {
	payload := gobEncode(getblocks{nodeAddress})
	request := append(commandToBytes("getblocks"), payload...)

	sendData(addr, request)
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

func sendData(addr string, data []byte) {
	conn, err := net.Dial(protocol, addr)
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

func sendGetData(address, kind string, id []byte) {
	payload := gobEncode(getdata{nodeAddress, kind, id})
	request := append(commandToBytes("getdata"), payload...)

	sendData(address, request)
}

func getTxHistory(b Block) {
	address := "15bcaRcuuxToXfPthPRVsXJhHHC42LLNuF" //FIXME; address has 2 badgercoin
	args := []string{"getwalletpubkeyhash", "-address", address}
	cmd := exec.Command("./blockchain", args...)

	pubKeyHash, err := cmd.CombinedOutput()
	if err != nil {
		return //FIXME
	}
	sPubKeyHash := string(pubKeyHash)

	for _, tx := range b.Transactions {
		for _, vin := range tx.Vin {
			sVinPK := string(vin.PubKey)
			if sVinPK == sPubKeyHash {
				fmt.Println("Equal!")
			}
		}
		for _, vout := range tx.Vout { //FIXME
			sVoutPK := string(vout.PubKeyHash)
			if sVoutPK == sPubKeyHash {
				fmt.Println("Equal!")
			}
		}
		//address in VIN: put to send TXs
		//address in VOUT and not in VIN: put to receive TXs
		//coinbase has no input
	}
}

func handleBlock(request []byte) {
	var buff bytes.Buffer
	var payload block

	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	blockData := payload.Block
	block := DeserializeBlock(blockData)
	getTxHistory(*block)

	if len(blocksInTransit) > 0 {
		blockHash := blocksInTransit[0]
		sendGetData(payload.AddrFrom, "block", blockHash)

		blocksInTransit = blocksInTransit[1:]
	}
}

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		fmt.Println("Block Deserialize error")
	}

	return &block
}

func handleInv(request []byte) {
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
		sendGetData(payload.AddrFrom, "block", blockHash)

		newInTransit := [][]byte{}
		for _, b := range blocksInTransit {
			if !bytes.Equal(b, blockHash) {
				newInTransit = append(newInTransit, b)
			}
		}
		blocksInTransit = newInTransit
	}
}

func handleConnection(conn net.Conn) {
	request, err := io.ReadAll(conn)
	if err != nil {
		log.Panic(err)
	}
	command := bytesToCommand(request[:commandLength])
	fmt.Printf("Received %s command\n", command)

	switch command {
	case "block":
		handleBlock(request)
	case "inv":
		handleInv(request)
	}
}

func startWalletServer() {
	ln, err := net.Listen(protocol, nodeAddress)
	if err != nil {
		log.Panic(err)
	}
	defer ln.Close()

	getBlocks()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Panic(err)
		}
		//not in goroutine as TX should be in order
		//OR get blocks with goroutines and sort it by timestamp
		handleConnection(conn)
	}
}
