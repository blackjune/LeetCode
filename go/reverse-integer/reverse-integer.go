package main

import (
	"fmt"
)

func reverse(x int) (r int) {
	if x > 2147483647 || x < -2147483647 {
		return 0
	}
	var nums []int
	for x != 0 {
		nums = append(nums, x%10)
		x /= 10
	}
	l := len(nums)
	for i := 0; i < l; i++ {
		if i == 0 && nums[i] == 0 {
			continue
		}
		r = r*10 + nums[i]
	}
	if r > 2147483647 || r < -2147483647 {
		return 0
	}
	return
}

func main() {
	fmt.Println(reverse(1534236469))
}
