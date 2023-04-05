package data

type Account struct {
	Balance float64
}

func CreateNewAccount(bal float64) Account {
	account := Account{
		Balance: bal,
	}
	return account
}
