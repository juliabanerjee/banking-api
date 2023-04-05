package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"banking-api/handlers"
)

// start with a basic GET endpoint and have sum hardcoded in
// need a router probably mux

// next step: take in user input and add/ subtract it from hardcoded balance
func main() {

	r := mux.NewRouter()

	fmt.Println("server is running")

	// need to connect to db and build a db with tableplus

	// GET
	r.HandleFunc("/account", handlers.GetTotalFunds).Methods("GET")
	// POST
	r.HandleFunc("/account", handlers.DepositFunds).Methods("POST")
	// PUT
	r.HandleFunc("/account", handlers.WithdrawFunds).Methods("PUT")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err)
	}
}
