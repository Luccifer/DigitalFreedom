package blockchain

import (
	"bytes"
	"crypto/sha512"
	"strconv"
)

func hashFor(block *Block) []byte {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))                                  // get the time and convert it into a unique series of digits
	headers := bytes.Join([][]byte{timestamp, block.PreviousBlockHash, block.AllData}, []byte{}) // concatenate all the block data
	hash := sha512.Sum512(headers)                                                               // hash the whole thing
	block.Hash = hash[:]
	return hash[:]
}

func hashString(s string) []byte {
	h := sha512.New()
	h.Write([]byte(s))
	hashed := h.Sum(nil)
	return hashed
}
