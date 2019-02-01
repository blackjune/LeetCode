package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	i := 0
	for i < len(nums) {
		if nums[i] == val {
			nums[i], nums[len(nums)-1] = nums[len(nums)-1], nums[i]
			nums = nums[:len(nums)-1]
		} else {
			i++
		}
	}
	return i
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	l := removeElement(nums, 2)
	fmt.Println(nums, l)
}
