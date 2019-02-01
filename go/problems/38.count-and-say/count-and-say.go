package main

import (
	"fmt"
)

func countAndSay(n int) (r string) {
	r = "1"
	for i := 1; i < n; i++ {
		r = newStr(r)
		fmt.Printf("%4d, %s\n", i+1, r)
	}
	return
}

func newStr(str string) string {
	sl := len(str)
	var prev string
	var count int = 0
	var i int
	var r string
	for i < sl {
		curr := str[i : i+1]
		if prev == "" || prev == curr {
			count++
		} else {
			r += fmt.Sprintf("%d%s", count, prev)
			count = 1
		}
		prev = curr
		i++
	}
	r += fmt.Sprintf("%d%s", count, prev)
	return r
}

func main() {
	fmt.Println(countAndSay(6))
}
