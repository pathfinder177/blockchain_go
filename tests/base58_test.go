package tests

import (
	"blockchain/common"
	"encoding/hex"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase58(t *testing.T) {
	rawHash := "00010966776006953D5567439E5E39F86A0D273BEED61967F6"
	hash, err := hex.DecodeString(rawHash)
	if err != nil {
		log.Fatal(err)
	}

	encoded := common.Base58Encode(hash)
	assert.Equal(t, "16UwLL9Risc3QfPqBUvKofHmBQ7wMtjvM", string(encoded))

	decoded := common.Base58Decode([]byte("16UwLL9Risc3QfPqBUvKofHmBQ7wMtjvM"))
	assert.Equal(t, strings.ToLower("00010966776006953D5567439E5E39F86A0D273BEED61967F6"), hex.EncodeToString(decoded))
}
