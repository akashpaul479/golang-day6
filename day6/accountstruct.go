package day6

import "fmt"

type account struct {
	Holder  string
	Balance float64
}

func (a *account) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("deposit amount must be positive")
		return
	}
	a.Balance += amount
	fmt.Println("Deposited", amount)

}

func (a *account) Withdrawl(amount float64) {
	if amount <= 0 {
		fmt.Println("withdrawl amount must be positive")
		return
	}
	if amount > a.Balance {
		fmt.Println("Insufficient balance")
		return
	}
	a.Balance -= amount
	fmt.Println("Withdrawn:", amount)
}

func Account() {
	acc := account{Holder: "akash", Balance: 1000}

	acc.Deposit(200)
	acc.Withdrawl(500)

	fmt.Println("Final balance:", acc.Balance)
}
