package blockchain

import (
	"bytes"
	math "math/rand"
	"strconv"
	"time"
)

var Pos = PoSNetwork{}

func init() {
	pushGenesisToHead()
}

func getGenesisBlock() *Block {
	math.Seed(time.Now().Unix())
	genesisTime := time.Now().Unix()

	timestamp := []byte(strconv.FormatInt(genesisTime, 10))
	data := "Genesis - The One"
	validator := "Void"
	hashedData := hashString(data)
	headers := bytes.Join([][]byte{timestamp, hashedData}, []byte{})
	genesisBlock := Block{genesisTime, []byte{}, headers, []byte(data), validator, []byte{}}

	return &genesisBlock
}

func pushGenesisToHead() {
	genesis := getGenesisBlock()
	Pos = PoSNetwork{
		Blockchain: []*Block{
			{
				genesis.Timestamp,
				genesis.PreviousBlockHash,
				genesis.Hash,
				genesis.AllData,
				genesis.Validator,
				genesis.Sign,
			},
		},
	}

	Pos.BlockchainHead = Pos.Blockchain[0]
}
