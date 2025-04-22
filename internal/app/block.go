package app

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"
)

type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
	Height        int
}

func NewBlock(transactions []*Transaction, prevBlockHash []byte, height int) *Block {
	block := &Block{time.Now().Unix(), transactions, prevBlockHash, []byte{}, 0, height}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func (b *Block) HashTransactions() []byte {
	var transactions [][]byte

	for _, tx := range b.Transactions {
		transactions = append(transactions, tx.Serialize())
	}
	mTree := NewMerkleTree(transactions)

	return mTree.RootNode.Data
}

func NewGenesisBlock(coinbase []*Transaction) *Block {
	return NewBlock(coinbase, []byte{}, 0)
}

func getCurrenciesNamesFromBlock(block *Block) map[string]struct{} {
	//map len depends on block size and TX handling options
	//here map has len/2 because each block has coinbase and 1 UTXO TX for any currency
	currenciesNames := make(map[string]struct{}, len(block.Transactions)/2)

	//set of currencies names. empty struct has size 0
	for _, tx := range block.Transactions {
		currenciesNames[tx.Currency] = struct{}{}
	}

	return currenciesNames
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		fmt.Println("Block Serialize error")
	}

	return result.Bytes()
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
