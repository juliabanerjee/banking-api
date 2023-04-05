package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"banking-api/handlers"
)

// start with a basic GET endpoint and have sum hardcoded in
// need a router probably mux
func main() {

	r := mux.NewRouter()

	fmt.Println("server is running")

	// GET
	r.HandleFunc("/account", handlers.GetTotalFunds).Methods("GET")
	// POST
	r.HandleFunc("/account", handlers.DepositFunds).Methods("POST")
	// PUT
	r.HandleFunc("/account", handlers.RemoveFunds).Methods("PUT")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err)
	}
}

// naming: bank account, adding money (DepositFunds/ Sum), subtracting money (RemoveFunds), getting money (GetTotalFunds)
