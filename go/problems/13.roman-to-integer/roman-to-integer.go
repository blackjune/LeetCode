package main

import (
	"fmt"
)

func one(s string) int {
	switch s {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	}
	return 0
}

func romanToInt(s string) int {
	if len(s) == 1 {
		return one(s)
	}
	var r int
	for i := 0; i < len(s); i += 1 {
		var next int
		curr := one(s[i : i+1])
		if i < len(s)-1 {
			next = one(s[i+1 : i+2])
		}
		if curr < next {
			r = r - curr
		} else {
			r = r + curr
		}
	}
	return r
}

func main() {
	fmt.Println(romanToInt("MMXCX"))
}
