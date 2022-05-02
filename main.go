package main

import (
	"fmt"

	"github.com/rosenrose/go-learn/accounts"
)

func main() {
	account := accounts.NewAccount("Hans")
	fmt.Println(*account)
}
