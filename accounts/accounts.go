package accounts

import (
	"errors"
	"fmt"
)

var errNoMoney = errors.New("cant't withdraw")

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount
func (a *Account) Deposit(amount int) {
	fmt.Println("Deposit", amount)
	a.balance += amount
}

// Balance of account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}

	fmt.Println("Withdraw", amount)
	a.balance -= amount

	return nil
}

// ChangeOwner
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.owner, "'s account. Has: ", a.balance)
}
