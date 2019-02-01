package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)
	l := len(str)
	for i := 0; i < l/2; i++ {
		if str[i] != str[l-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(10))
}
