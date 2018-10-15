package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	Items []string
}

func NewStack() *Stack {
	return &Stack{
		Items: make([]string, 0),
	}
}
func (s *Stack) Push(str string) {
	s.Items = append(s.Items, str)
}

func (s *Stack) Pop() (string, error) {
	itemsCount := len(s.Items)
	if itemsCount == 0 {
		return "", errors.New("stack empty")
	}
	r := s.Items[itemsCount-1]
	s.Items = s.Items[0 : itemsCount-1]
	return r, nil
}

func (s *Stack) Peek() string {
	l := len(s.Items)
	if l == 0 {
		return ""
	}
	return s.Items[l-1]
}

func (s *Stack) Len() int {
	return len(s.Items)
}

func isValid(s string) bool {
	sLen := len(s)
	if sLen == 0 {
		return true
	}
	if len(s) < 2 {
		return false
	}
	stack := NewStack()
	match := map[string]string{
		"(": ")",
		")": "(",
		"{": "}",
		"}": "{",
		"[": "]",
		"]": "[",
	}
	for i := 0; i < sLen; i++ {
		curr := s[i : i+1]
		if curr == "(" || curr == "{" || curr == "[" {
			stack.Push(curr)
		} else if m, ok := match[curr]; ok {
			prev, err := stack.Pop()
			if err != nil || m != prev {
				return false
			}
		}
	}
	return stack.Len() == 0
}

func main() {
	fmt.Println(isValid("{()}"))
}
