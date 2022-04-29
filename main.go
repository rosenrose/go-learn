package main

import (
	"fmt"
)

func main() {
	fmt.Println(repeat(15, 13, 11, 9, 7))
}

func repeat(numbers ...int) int {
	total := 0

	for index, number := range numbers {
		fmt.Println(index, number)
		total += number
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}

	return total
}
