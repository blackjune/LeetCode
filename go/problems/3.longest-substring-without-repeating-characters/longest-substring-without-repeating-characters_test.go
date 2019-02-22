package characters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	dataSet := [][]interface{}{
		{"dcdfdfccdssgag", 3},
	}
	for _, data := range dataSet {
		n := data[0].(string)
		m := data[1].(int)
		assert.Equal(t, m, lengthOfLongestSubstring(n))
	}
}
