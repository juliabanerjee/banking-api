package handlers

import (
	"fmt"
	"net/http"
)

func GetTotalFunds(w http.ResponseWriter, r *http.Request) {
	balance := 200.00
	fmt.Printf("Your balance is: £%f \n", balance)

}

func DepositFunds(w http.ResponseWriter, r *http.Request) {
	addBal := 250.00
	fmt.Printf("Your new balance is: £%f \n", addBal)

}

func RemoveFunds(w http.ResponseWriter, r *http.Request) {
	subBal := 150.00
	fmt.Printf("Your new balance is: £%f \n", subBal)
}
