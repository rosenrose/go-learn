package main

import (
	"fmt"

	"github.com/rosenrose/go-learn/accounts"
)

func main() {
	account := accounts.NewAccount("Hans")
	fmt.Println(*account)

	account.Deposit(300)
	fmt.Println(account.Balance())
	
	err := account.Withdraw(500)
	if err != nil {
		// log.Fatalln(err)
		fmt.Println(err)
	}
	err = account.Withdraw(200)
	
	account.ChangeOwner("Alice")
	fmt.Println(account.Owner(), account.Balance(), err)
}
