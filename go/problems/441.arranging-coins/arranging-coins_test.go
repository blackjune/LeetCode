package coins

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	dataSet := [][]interface{}{
		{5, 2},
		{6, 3},
		{8, 3},
		{1, 1},
		{2, 1},
	}
	for _, data := range dataSet {
		nums1 := data[0].(int)
		m := data[1].(int)
		assert.Equal(t, m, arrangeCoins(nums1))
	}
}
