package api

import (
	"DigitalFreedom/blockchain"
	"fmt"
	math "math/rand"
	"net/http"
	"time"
)

func addNew(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		for i, block := range blockchain.Pos.Blockchain { // iterate on each block
			fmt.Fprintf(w, "\n------------------------------------------------\n")
			fmt.Fprintf(w, "Block ID : %d \n", i)
			fmt.Fprintf(w, "Timestamp : %d \n", block.Timestamp+int64(i))
			fmt.Fprintf(w, "Hash of the block : %x\n", block.MyBlockHash)
			fmt.Fprintf(w, "Hash of the previous Block : %x\n", block.PreviousBlockHash)
			fmt.Fprintf(w, "All the transactions : %s\n", block.AllData)
			fmt.Fprintf(w, "\n------------------------------------------------\n")
		}
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		mu.Lock()
		math.Seed(time.Now().UnixNano())

		//blockchain.Chain.AddBlock(r.FormValue("data"))
		mu.Unlock()

		for i, block := range blockchain.Pos.Blockchain { // iterate on each block
			fmt.Fprintf(w, "\n------------------------------------------------\n")
			fmt.Fprintf(w, "Block ID : %d \n", i)
			fmt.Fprintf(w, "Timestamp : %d \n", block.Timestamp+int64(i))
			fmt.Fprintf(w, "Hash of the block : %x\n", block.MyBlockHash)
			fmt.Fprintf(w, "Hash of the previous Block : %x\n", block.PreviousBlockHash)
			fmt.Fprintf(w, "All the transactions : %s\n", block.AllData)
			fmt.Fprintf(w, "\n------------------------------------------------\n")
		}

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
