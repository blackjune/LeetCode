package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	dataSet := [][]interface{}{
		{2, 1},
		{4, 3},
		{5, 5},
		{6, 8},
	}
	for _, data := range dataSet {
		n := data[0].(int)
		m := data[1].(int)
		assert.Equal(t, fib(n), m)
	}
}
