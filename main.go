package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(multiply(2, 3))

	length, upper := lenAndUpper("haha")
	length2, _ := lenAndUpperNaked("vobobo")

	fmt.Println(length, upper, length2)

	repeat("a", "b", "C", "dfsd af")
}

func multiply(a, b int) int {
	defer fmt.Println("done")
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeat(words ...string) {
	fmt.Println(words)
}

func lenAndUpperNaked(name string) (length int, upper string) {
	length = len(name)
	upper = strings.ToUpper(name)
	return
}
