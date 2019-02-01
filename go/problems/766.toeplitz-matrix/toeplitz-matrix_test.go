package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	dataSet := [][]interface{}{
		{[][]int{
			{1, 2, 3, 4},
			{5, 1, 2, 3},
			{9, 5, 1, 2},
		}, true},
	}
	for _, data := range dataSet {
		nums1 := data[0].([][]int)
		nums2 := data[1].(bool)
		assert.Equal(t, isToeplitzMatrix(nums1), nums2)
	}
}
