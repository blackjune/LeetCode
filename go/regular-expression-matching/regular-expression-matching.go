package main

import (
	"fmt"
	"strings"
)

type Pattern struct {
	expect string
	min    int
	max    int
}

func parsePattern(p string) []*Pattern {
	var matches []*Pattern
	l := len(p)
	j := 0
	var isNew bool = true

	for i := 0; i < l; i++ {
		expect := p[i : i+1]
		max := 1
		min := 1
		if i > 0 && p[i-1:i] == "*" {
			isNew = true
		}
		if i < l-1 && p[i+1:i+2] == "*" {
			max = -1
			min = 0
			i++
			isNew = true
		}
		if isNew {
			matches = append(matches, &Pattern{
				expect: expect,
				min:    min,
				max:    max,
			})
			j++
		} else if j > 0 {
			matches[j-1].expect += expect
			matches[j-1].min = 1
			matches[j-1].max = 1
		}
		fmt.Println(expect)
		if expect != "." {
			isNew = false
		}
	}
	return matches
}

func isMatch(s string, p string) bool {
	patterns := parsePattern(p)
	for _, i := range patterns {
		fmt.Println(i)
	}
	return match(s, patterns)
}

func match(s string, patterns []*Pattern) bool {
	patternIndex := 0
	i := 0
	strLen := len(s)
	for patternIndex < len(patterns) {
		pattern := patterns[patternIndex]
		if pattern.max == 1 {
			expectLen := len(pattern.expect)
			if len(pattern.expect) > strLen {
				fmt.Println(1)
				return false
			}
			for ; i < expectLen; i++ {
				if s[i] != pattern.expect[i] && pattern.expect[i:i+1] != "." {
					fmt.Println(2)
					return false
				}
			}
		} else if pattern.max == -1 {
			flg, step := matchStar(s[i:], patterns[patternIndex:])
			if flg == false {
				fmt.Println(3)
				return false
			}
			patternIndex += step
		}
		patternIndex++
	}
	return true
}

func matchStar(s string, patterns []*Pattern) (bool, int) {
	check := ""
	patternIndex := 0
	for check == "" && patternIndex < len(patterns) {
		pattern := patterns[patternIndex]
		fmt.Println(pattern)
		if pattern.max != -1 {
			check = pattern.expect
			break
		}
		patternIndex++
	}
	if check == "." {
		return true, patternIndex
	}
	pos := strings.Index(s, check)
	fmt.Printf("s: %s, check: %s, pos: %d\n", s, check, pos)
	if pos < 0 {
		return false, patternIndex
	}
	newS := s[0 : pos+1]
	
	return true, patternIndex
}

func main() {
	p := "mis*is*p*."
	fmt.Println(isMatch("mississippi", p))
}
