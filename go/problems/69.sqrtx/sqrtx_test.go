package sqrtx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMySqrt(t *testing.T) {
	dataSet := [][]interface{}{
		{4, 2},
		{9, 3},
		{8, 2},
		{16, 4},
		{1, 1},
		{2, 1},
	}
	for _, data := range dataSet {
		assert.Equal(t, data[1].(int), mySqrt(data[0].(int)))
	}
}
