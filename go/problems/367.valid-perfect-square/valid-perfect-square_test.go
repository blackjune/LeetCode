package square

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	dataSet := [][]interface{}{
		{1, true},
		{16, true},
		{24, false},
		{25, true},
	}
	for _, data := range dataSet {
		nums1 := data[0].(int)
		m := data[1].(bool)
		assert.Equal(t, m, isPerfectSquare(nums1))
	}
}
