package main

import (
	"fmt"
)

func main() {
	fmt.Println(isDrink(17))
}

func isDrink(age int) bool {
	if kAge := age + 1; kAge < 18 {
		return false
	}
	return true
}
