package common

import (
	"bytes"
	"encoding/binary"
	"log"
)

// IntToHex converts an int64 to a byte array
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// ReverseBytes reverses a byte array
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

// func wordsToBytes(words []big.Word) []byte {
// 	bytes := make([]byte, len(words)*8) // Each big.Word is 8 bytes on 64-bit systems
// 	for i, word := range words {
// 		binary.BigEndian.PutUint64(bytes[i*8:(i+1)*8], uint64(word))
// 	}
// 	return bytes
// }
