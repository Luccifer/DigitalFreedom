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

	http.HandleFunc("/block", addNew)
	http.HandleFunc("/node", addNewNode)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
