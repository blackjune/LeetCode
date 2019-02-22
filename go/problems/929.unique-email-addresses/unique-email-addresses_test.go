package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	dataSet := [][]interface{}{
		{[]int{3, 1, 2, 4}, []int{2, 4, 1, 3}},
		{[]int{3, 2, 1, 4}, []int{2, 4, 1, 3}},
	}
	for _, data := range dataSet {
		nums1 := data[0].([]int)
		nums2 := data[1].([]int)
		assert.Equal(t, sortArrayByParity(nums1), nums2)
	}
}
