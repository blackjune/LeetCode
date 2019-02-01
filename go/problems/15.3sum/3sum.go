package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	var had = make(map[string]struct{})
	comb(len(nums), 3, func(c []int) {
		if nums[c[0]]+nums[c[1]]+nums[c[2]] == 0 {
			tmp := []int{nums[c[0]], nums[c[1]], nums[c[2]]}
			sort.Sort(sort.IntSlice(tmp))
			key := fmt.Sprint(tmp)
			if _, ok := had[key]; !ok {
				result = append(result, tmp)
				had[key] = struct{}{}
			}
		}
	})
	return result
}

func comb(n, m int, emit func([]int)) {
	s := make([]int, m)
	last := m - 1
	var rc func(int, int)
	rc = func(i, next int) {
		for j := next; j < n; j++ {
			s[i] = j
			if i == last {
				emit(s)
			} else {
				rc(i+1, j+1)
			}
		}
		return
	}
	rc(0, 0)
}

func main() {

	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
