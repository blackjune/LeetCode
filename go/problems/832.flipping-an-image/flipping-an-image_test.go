package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	dataSet := [][]interface{}{
		{[][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}, [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}}},
	}
	for _, data := range dataSet {
		nums1 := data[0].([][]int)
		nums2 := data[1].([][]int)
		assert.Equal(t, flipAndInvertImage(nums1), nums2)
	}
}
