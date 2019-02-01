package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedSquares(t *testing.T) {
	dataSet := [][]interface{}{
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	}
	for _, data := range dataSet {
		nums1 := data[0].([]int)
		nums2 := data[1].([]int)
		res := sortedSquares(nums1)
		assert.Equal(t, res, nums2)
	}
}
