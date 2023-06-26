package gopointers

import (
	"errors"
	"fmt"
)

var ErrInsufficentFunds = errors.New("insufficent funds, can't withdraw")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrInsufficentFunds
	}
	w.balance -= amount
	return nil

}

func (w *Wallet) Balance() Bitcoin {
	//fmt.Printf("Address of balance in wallet: %v\n", &w.balance)
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
