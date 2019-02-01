package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	dataSet := [][]interface{}{
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
	}
	for _, data := range dataSet {
		nums1 := data[0].([]int)
		m := data[1].(int)
		assert.Equal(t, m, minCostClimbingStairs(nums1))
	}
}
