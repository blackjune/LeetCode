package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	dataSet := [][]interface{}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{0}, 0, []int{1}, 1, []int{1}},
		{[]int{2, 5, 6, 0, 0, 0}, 3, []int{1, 2, 4}, 3, []int{1, 2, 2, 4, 5, 6}},
	}
	for _, data := range dataSet {
		nums1 := data[0].([]int)
		m := data[1].(int)
		nums2 := data[2].([]int)
		n := data[3].(int)
		merge(nums1, m, nums2, n)
		assert.Equal(t, data[4].([]int), nums1)
	}
}
