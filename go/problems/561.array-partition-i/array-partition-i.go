package array

import (
	"sort"
)

func arrayPairSum(nums []int) int {
	sort.Sort(sort.IntSlice(nums))
	sum := 0
	for i := 0; i < len(nums)-1; i += 2 {
		sum += nums[i]
	}
	return sum
}
