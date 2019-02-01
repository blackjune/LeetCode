package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums)-1; {
		if nums[i] == nums[i+1] {
			nums = append(nums[0:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)
}

func main() {
	nums := []int{1, 1}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
