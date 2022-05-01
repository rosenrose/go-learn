package main

import (
	"fmt"
)

func main() {
	names := [5]string {"a", "b", "C"}
	names[3] = "asd f"
	names[4] = "zx cv"

	nums := []int {1, 2, 3}
	nums = append(nums, 4)
	fmt.Println(names, nums)
}
