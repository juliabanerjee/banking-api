package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"banking-api/data"
)

// things to consider:

// the sum will be a float but it will be returned in a string format so there may need to be some type conversion. JSON accepts floats but not ints/ floats written as strings

// need it to take int and either add or subtract it to balance
// sprintf for a simpler way
type Deposit struct {
	Amount float64
}
type Withdrawal struct {
	Amount float64
}

func GetTotalFunds(w http.ResponseWriter, r *http.Request) {

	newAcc := data.CreateNewAccount(255.00)
	// converting it right now because it is hardcoded and not being sent through JSON
	strBalance := "£" + strconv.FormatFloat(newAcc.Balance, 'f', 2, 64)
	printStmnt := "Your balance is: " + strBalance
	io.WriteString(w, printStmnt)

}

func DepositFunds(w http.ResponseWriter, r *http.Request) {
	newAcc := data.CreateNewAccount(255.00)
	var deposit Deposit
	// read from request body and use that to add onto account balance

	err := json.NewDecoder(r.Body).Decode(&deposit)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	// new total for the moment but this should be handled by GET endpoint, addition here >>

	newAcc.Balance = newAcc.Balance + deposit.Amount

	newBalStr := strconv.FormatFloat(newAcc.Balance, 'f', 2, 64)
	printStmnt := "Your new balance is: " + newBalStr

	io.WriteString(w, printStmnt)
}

func WithdrawFunds(w http.ResponseWriter, r *http.Request) {
	// create account and withdrawal struct
	newAcc := data.CreateNewAccount(255.00)
	var withdrawal Withdrawal
	// decode amount in request body into withdrawal struct
	err := json.NewDecoder(r.Body).Decode(&withdrawal)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	// subtract withdrawl from account balance
	newAcc.Balance = newAcc.Balance - withdrawal.Amount
	// convert new account balance into string so it can be printed
	newBalStr := strconv.FormatFloat(newAcc.Balance, 'f', 2, 64)
	//fmt.Printf("Your new balance is: %v \n", subBal)
	printStmnt := "Your balance is: " + newBalStr

	io.WriteString(w, printStmnt)
}

// NB: find way to not use io.WriteString so you can avoid type conversion!!

// current issues: the api only does the subtraction once and will not continue to subtract when multiple of the same request is made. if you change the amount to be subtracted then it will still subtract from original balance rather than the updated balance
