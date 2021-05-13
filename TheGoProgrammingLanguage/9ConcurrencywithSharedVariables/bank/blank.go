package bank

var balance int

// Deposit,存款
func Deposit(amount int) {
	balance += amount
}

func Balance() int {
	return balance
}