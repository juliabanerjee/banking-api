package handlers

import (
	"io"
	"net/http"
)

// things to consider:

// the sum will be a float but it will be returned in a string format so there may need to be some type conversion

func GetTotalFunds(w http.ResponseWriter, r *http.Request) {
	balance := "£200.00"
	printStmnt := "Your balance is:" + balance
	io.WriteString(w, printStmnt)

}

func DepositFunds(w http.ResponseWriter, r *http.Request) {
	addBal := "£250.00"
	printStmnt := "Your balance is:" + addBal

	io.WriteString(w, printStmnt)
}

func RemoveFunds(w http.ResponseWriter, r *http.Request) {
	subBal := "£150.00"
	//fmt.Printf("Your new balance is: %v \n", subBal)
	printStmnt := "Your balance is:" + subBal

	io.WriteString(w, printStmnt)
}
