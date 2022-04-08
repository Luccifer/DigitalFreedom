package api

import (
	"DigitalFreedom/blockchain"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var pos = &blockchain.Pos

func HandleWebServer() {

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/block", getBlock)

	http.HandleFunc("/node", nodeHandler)
	http.HandleFunc("/id", personHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
