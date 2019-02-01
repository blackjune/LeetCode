package main

import (
	"fmt"
)

func maxArea(height []int) int {
	var max int
	var curr int
	for i := 0; i < len(height)-1; i++ {
		for j := 1; j < len(height); j++ {
			if height[i] > height[j] {
				curr = height[j] * (j - i)
			} else {
				curr = height[i] * (j - i)
			}
			if curr > max {
				max = curr
			}
		}
	}
	return max
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
