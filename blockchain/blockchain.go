package blockchain

import (
	"bytes"
	"crypto/sha512"
	math "math/rand"
	"strconv"
	"time"
)

var Pos = PoSNetwork{}

func init() {

	math.Seed(time.Now().Unix())

	// generate an initial PoS network including a blockchain with a genesis block.
	genesisTime := time.Now().Unix()
	genesisString := string(genesisTime)

	h := sha512.New()
	h.Write([]byte("GB547"))
	hashed := h.Sum(nil)

	dataH := sha512.New()
	dataH.Write([]byte("Genesis Block"))
	dataHashed := dataH.Sum(nil)

	timestamp := []byte(strconv.FormatInt(genesisTime, 10)) // get the time and convert it into a unique series of digits
	headers := bytes.Join([][]byte{timestamp}, hashed)      // concatenate all the block data
	hash := sha512.Sum512(headers)                          // hash the whole thing

	Pos = PoSNetwork{
		Blockchain: []*Block{
			{
				Timestamp:         genesisTime,
				PreviousBlockHash: hash[:],
				Hash:              newHash(genesisString),
				AllData:           dataHashed,
				Validator:         "GB547",
			},
		},
	}
	Pos.BlockchainHead = Pos.Blockchain[0]

}

//method for generating a hash of the block
//concatenate all the data and hash it to obtain the block hash
func (block *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))                                  // get the time and convert it into a unique series of digits
	headers := bytes.Join([][]byte{timestamp, block.PreviousBlockHash, block.AllData}, []byte{}) // concatenate all the block data
	hash := sha512.Sum512(headers)                                                               // hash the whole thing
	block.Hash = hash[:]                                                                         // now set the hash of the block
}

func NewBlockHash(block *Block) []byte {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))                                  // get the time and convert it into a unique series of digits
	headers := bytes.Join([][]byte{timestamp, block.PreviousBlockHash, block.AllData}, []byte{}) // concatenate all the block data
	hash := sha512.Sum512(headers)                                                               // hash the whole thing
	block.Hash = hash[:]
	return hash[:]
}

func newHash(s string) []byte {
	h := sha512.New()
	h.Write([]byte(s))
	hashed := h.Sum(nil)
	return hashed
	//return hex.EncodeToString(hashed)
}
