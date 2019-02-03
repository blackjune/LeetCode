package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	dataSet := [][]interface{}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	}
	for _, data := range dataSet {
		nums1 := data[0].([]int)
		m := data[1].(int)
		assert.Equal(t, m, maxSubArray(nums1))
	}
}
