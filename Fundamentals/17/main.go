package main

import "fmt"

type BankAccount struct {
	name   string
	amount int
}

func NewAccount() *BankAccount {
	return &BankAccount{name: "", amount: 0}
}

func (account *BankAccount) debit(value int) int {

	account.amount -= value
	return account.amount
}

func main() {

	client := BankAccount{
		name:   "Augusto",
		amount: 100,
	}

	newUser := NewAccount()
	fmt.Println(newUser)
	newUser.amount = 10
	newUser.name = "João"
	fmt.Println(newUser)

	fmt.Println(client.debit(50))
	fmt.Println(client.amount)
}
