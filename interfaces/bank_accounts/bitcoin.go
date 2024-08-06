package main

import "errors"

type Bitcoin struct {
	name    string
	balance int
	fee     int
}

func NewBitcoin() *Bitcoin {
	return &Bitcoin{
		name:    "Bitcoin",
		balance: 0,
		fee:     3,
	}
}

func (b *Bitcoin) GetName() string {
	return b.name
}

func (b *Bitcoin) GetBalance() int {
	return b.balance
}

func (b *Bitcoin) Deposit(amount int) {
	b.balance += amount
}

func (b *Bitcoin) Withdraw(amount int) error {
	newBalance := b.balance - amount - b.fee
	if newBalance < 0 {
		return errors.New("insufficient funds")
	}
	b.balance = newBalance
	return nil
}
