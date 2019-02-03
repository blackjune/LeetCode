package dynamic_programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	dataSet := [][]interface{}{
		{2, 2},
		{3, 3},
		{4, 5},
	}
	for _, data := range dataSet {
		n := data[0].(int)
		m := data[1].(int)
		assert.Equal(t, m, climbStairs(n))
	}
}
