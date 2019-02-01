package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	dataSet := [][]interface{}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}
	for _, data := range dataSet {
		nums1 := data[0].([]int)
		m := data[1].(int)
		assert.Equal(t, m, maxProfit(nums1))
	}
}
