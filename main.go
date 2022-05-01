package main

import (
	"fmt"
)

func main() {
	a := 1
	b := &a
	a = 6
	fmt.Println(a, *b)
	
	*b = 3
	fmt.Println(a, *b)

	fmt.Println(a, &a, b, *b, &b)
}
