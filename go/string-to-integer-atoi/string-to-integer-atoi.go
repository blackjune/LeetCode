package main

import (
	"fmt"
	"math"
	"strings"
)

func myAtoi(str string) int {
	str = strings.Trim(str, " ")
	digitFlg := false
	overFlag := false
	sign := ""
	intMap := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"0": 0,
	}
	var res int
	for _, r := range str {
		s := string(r)
		if (s == "+" || s == "-") && digitFlg == false && sign == "" {
			sign = s
		} else if num, ok := intMap[string(r)]; ok {
			if res+1 < res || res+1 > math.MaxInt32 {
				overFlag = true
				break
			}
			res = res*10 + num
			if res+1 > math.MaxInt32+1 {
				overFlag = true
				break
			}
			digitFlg = true
		} else {
			break
		}
	}
	if overFlag {
		if sign == "-" {
			return math.MinInt32
		}
		return math.MaxInt32
	}
	if sign == "-" {
		return -res
	}
	return res
}

func main() {
	fmt.Println(myAtoi("2147483648"))
}
