package api

import (
	"DigitalFreedom/blockchain"
	"fmt"
	"net/http"
	"strconv"
)

func nodeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		switch len(pos.Validators) {
		case 0:
			fmt.Fprintf(w, "\n-----------------NO VALIDATORS--------------------\n")
		default:
			for _, validator := range pos.Validators { // iterate on each block
				fmt.Fprintf(w, "\n-----------------VALIDATORS--------------------\n")
				fmt.Fprintf(w, "Validator Adderess : %d \n", validator.Address)
				fmt.Fprintf(w, "Stake : %d", validator.Stake)
				fmt.Fprintf(w, "\n------------------------------------------------\n")
			}
		}
	case "POST":
		digits, _ := strconv.ParseInt(r.FormValue("stake"), 10, 64)
		address := r.FormValue("address")

		blockchain.NewNode(int(digits), address)
		fmt.Fprintf(w, "\n-----------------VALIDATOR ADDED----------------\n")
		fmt.Fprintf(w, "Validator Adderess : %d \n", address)
		fmt.Fprintf(w, "Stake : %d", digits)
		fmt.Fprintf(w, "\n------------------------------------------------\n")
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
