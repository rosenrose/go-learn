package main

import (
	"fmt"
)

type person struct {
	name string
	age int
	food []string
}

func main() {
	man := person {"mike", 20, []string {"a", "b"}}
	woman := person {age: 25, name: "yolo"}
	a := person {}

	fmt.Println(man, man.name, woman)
	fmt.Println(a)
}
