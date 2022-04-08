package blockchain

import (
	"bytes"
	"errors"
	"fmt"
	math "math/rand"
	"time"
)

type PoSNetwork struct {
	Blockchain     []*Block
	BlockchainHead *Block
	Validators     []*Node
}

func init() {
	print(randAddress())
}

func GenerateNewBlock(Validator *Node) ([]*Block, *Block, error) {
	if err := ValidateBlockchain(); err != nil {
		Validator.Stake -= 10
		return Pos.Blockchain, Pos.BlockchainHead, err
	}

	currentTime := time.Now().Unix()

	newBlock := &Block{
		Timestamp:         currentTime,
		PreviousBlockHash: Pos.BlockchainHead.Hash,
		Hash:              hashFor(Pos.BlockchainHead),
		Validator:         Validator.Address,
	}

	if err := ValidateBlockCandidate(newBlock); err != nil {
		Validator.Stake -= 10
		return Pos.Blockchain, Pos.BlockchainHead, err
	} else {
		Pos.Blockchain = append(Pos.Blockchain, newBlock)
	}
	return Pos.Blockchain, newBlock, nil
}

func ValidateBlockchain() error {
	if len(Pos.Blockchain) <= 1 {
		return nil
	}

	currBlockIdx := len(Pos.Blockchain) - 1
	prevBlockIdx := len(Pos.Blockchain) - 2

	for prevBlockIdx >= 0 {
		currBlock := Pos.Blockchain[currBlockIdx]
		prevBlock := Pos.Blockchain[prevBlockIdx]

		hashCheck := bytes.Compare(currBlock.PreviousBlockHash, prevBlock.Hash)
		if hashCheck != 0 {
			return errors.New("blockchain has inconsistent hashes")
		}

		if currBlock.Timestamp <= prevBlock.Timestamp {
			return errors.New("blockchain has inconsistent timestamps")
		}

		hashCheck2 := bytes.Compare(hashFor(prevBlock), currBlock.Hash)
		if hashCheck2 != 0 {
			return errors.New("blockchain has inconsistent hash generation")
		}
		currBlockIdx--
		prevBlockIdx--
	}
	return nil
}

func ValidateBlockCandidate(newBlock *Block) error {

	hashCheck := bytes.Compare(Pos.BlockchainHead.Hash, newBlock.PreviousBlockHash)
	if hashCheck != 0 {
		return errors.New("blockchain HEAD hash is not equal to new block previous hash")
	}

	if Pos.BlockchainHead.Timestamp >= newBlock.Timestamp {
		return errors.New("blockchain HEAD timestamp is greater than or equal to new block timestamp")
	}

	hashCheck2 := bytes.Compare(hashFor(Pos.BlockchainHead), newBlock.Hash)
	if hashCheck2 != 0 {
		return errors.New("new block hash of blockchain HEAD does not equal new block hash")
	}
	return nil
}

func NewNode(stake int, validator string) []*Node {
	newNode := &Node{
		Stake:   stake,
		Address: validator,
	}
	Pos.Validators = append(Pos.Validators, newNode)
	return Pos.Validators
}

func randAddress() string {
	b := make([]byte, 16)
	_, _ = math.Read(b)
	return fmt.Sprintf("%x", b)
}

func SelectValidator() (*Node, error) {
	var winnerPool []*Node
	totalStake := 0
	for _, node := range Pos.Validators {
		if node.Stake > 0 {
			winnerPool = append(winnerPool, node)
			totalStake += node.Stake
		}
	}
	if winnerPool == nil {
		return nil, errors.New("there are no nodes with stake in the network")
	}
	winnerNumber := math.Intn(totalStake)
	tmp := 0
	for _, node := range Pos.Validators {
		tmp += node.Stake
		if winnerNumber < tmp {
			return node, nil
		}
	}
	return nil, errors.New("a winner should have been picked but wasn't")
}
