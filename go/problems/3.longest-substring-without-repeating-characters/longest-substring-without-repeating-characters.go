package characters

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	var start, max, count int
	have := make(map[rune]struct{})
	for _, c := range s {
		fmt.Printf("%v, count %d max %d\n", have, count, max)
		if _, ok := have[c]; ok {
			for k := range have {
				delete(have, k)
			}
			if count > max {
				max = count
			}
			start++
			count = 0
		} else {
			have[c] = struct{}{}
			count++
		}
	}
	if count > max {
		return count
	}
	return max
}
