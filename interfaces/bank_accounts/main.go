package main

import "fmt"

// BankAccount is an interface that defines the methods that a bank account should have.
type BankAccount interface {
	// GetName returns the name of the account.
	GetName() string
	// GetBalance returns the current balance of the account.
	GetBalance() int
	// Deposit adds the specified amount to the account.
	Deposit(amount int)
	// Withdraw subtracts the specified amount from the account. Could return an error if the account does not have enough funds.
	Withdraw(amount int) error
}

func main() {
	bankAccounts := []BankAccount{
		NewWellsFargo(),
		NewBitcoin(),
	}

	for _, account := range bankAccounts {
		accountName := account.GetName()
		fmt.Printf("Working with %s account\n", accountName)

		currentBalance := account.GetBalance()
		fmt.Printf("Starting balance: %d\n", currentBalance)

		account.Deposit(200)
		err := account.Withdraw(50)
		if err != nil {
			fmt.Printf("Could not complete withdrawl: %s\n", err)
		}

		currentBalance = account.GetBalance()
		fmt.Printf("Ending balance: %d\n", currentBalance)

		fmt.Println()
	}
}
