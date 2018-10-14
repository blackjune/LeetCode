package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type NumSlice []string

func (p NumSlice) Len() int {
	return len(p)
}
func (p NumSlice) Less(i, j int) bool {
	a := p[i] + p[j]
	b := p[j] + p[i]
	return a < b
}
func (p NumSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return "0"
	}
	if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	}
	numStrs := make([]string, len(nums))
	for i, num := range nums {
		numStrs[i] = strconv.Itoa(num)
	}
	sort.Sort(sort.Reverse(NumSlice(numStrs)))
	if numStrs[0] == "0" {
		return "0"
	}
	return strings.Join(numStrs, "")
}

func main() {
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
}
