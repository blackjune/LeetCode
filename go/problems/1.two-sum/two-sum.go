package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	l := len(nums)
	for i := 0; i < l; i++ {
		if j, ok := dict[target-nums[i]]; ok && j != i {
			return []int{j, i}
		}
		dict[nums[i]] = i
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
