package main

import (
	"fmt"

	"github.com/rosenrose/go-learn/accounts"
)

func main() {
	account := accounts.NewAccount("Hans")
	account.Deposit(300)
	fmt.Println(account.Balance())
	
	err := account.Withdraw(500)
	if err != nil {
		// log.Fatalln(err)
		fmt.Println(err)
	}
	err = account.Withdraw(200)
	fmt.Println(account.Balance(), err)
}
