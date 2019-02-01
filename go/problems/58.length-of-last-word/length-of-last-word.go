package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	l := 0
	for j := len(s) - 1; j >= 0; j-- {
		if s[j] == ' ' {
			if l > 0 {
				break
			} else {
				continue
			}
		} else {
			l++
		}
	}
	return l
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World   cfl"))
}
