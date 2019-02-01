package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	dataSet := [][]interface{}{
		{[]int{1, 4, 3, 2}, 4},
	}
	for _, data := range dataSet {
		nums1 := data[0].([]int)
		m := data[1].(int)
		assert.Equal(t, arrayPairSum(nums1), m)
	}
}
