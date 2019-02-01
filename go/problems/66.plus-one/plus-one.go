package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	n := 1
	var r []int
	for j := len(digits) - 1; j >= 0; j-- {
		k := n + digits[j]
		r = append([]int{k % 10}, r...)
		n = k / 10
	}
	if n > 0 {
		r = append([]int{n}, r...)
	}
	return r
}

func main() {
	fmt.Println(plusOne([]int{1, 9, 9, 9}))
}
