package main

import (
	"fmt"
)

func main() {
	charMap := map[string]int {"a": 123, "b c": 999, "d:f": 3}

	fmt.Println(charMap)

	for key, val := range charMap {
		fmt.Println(key, val)
	}
}
