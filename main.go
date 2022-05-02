package main

import (
	"fmt"

	"github.com/rosenrose/go-learn/banking"
)

func main() {
	account := banking.Account {Owner: "Hans", Balance: 300}
	fmt.Println(account)
}
