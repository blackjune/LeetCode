package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	dataSet := [][]interface{}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}},
	}
	for _, data := range dataSet {
		nums1 := data[0].([][]int)
		nums2 := data[1].([][]int)
		assert.Equal(t, transpose(nums1), nums2)
	}
}
