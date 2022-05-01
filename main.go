package main

import (
	"fmt"
)

func main() {
	fmt.Println(isDrink(18))
}

func isDrink(age int) bool {
	switch kAge := age + 1; kAge {
	case 10:
		return false
	case 18:
		return true
	}

	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 70:
		return false
	}
	
	return false
}
