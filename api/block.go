package api

import (
	"DigitalFreedom/blockchain"
	"fmt"
	"net/http"
)

func getBlock(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		for i, block := range blockchain.Pos.Blockchain { // iterate on each block
			fmt.Fprintf(w, "\n------------------------------------------------\n")
			fmt.Fprintf(w, "Block ID : %d \n", i)
			fmt.Fprintf(w, "Timestamp : %d \n", block.Timestamp+int64(i))
			fmt.Fprintf(w, "Hash of the block : %x\n", block.Hash)
			fmt.Fprintf(w, "Hash of the previous Block : %x\n", block.PreviousBlockHash)
			fmt.Fprintf(w, "All the transactions : %s\n", block.AllData)
			fmt.Fprintf(w, "\n------------------------------------------------\n")
		}
	default:
		fmt.Fprintf(w, "Sorry, only GET methods are supported.")
	}
}
