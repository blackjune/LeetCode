package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	dataSet := [][]interface{}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 7, 11, 15}, 18, []int{2, 3}},
		{[]int{2, 7, 11, 15}, 26, []int{3, 4}},
		{[]int{2, 7, 11, 15}, 13, []int{1, 3}},
	}
	for _, data := range dataSet {
		nums1 := data[0].([]int)
		m := data[1].(int)
		nums2 := data[2].([]int)
		assert.Equal(t, twoSum(nums1, m), nums2)
	}
}
