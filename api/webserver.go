package api

import (
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex

func HandleWebServer() {

	http.HandleFunc("/block", addNew)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
