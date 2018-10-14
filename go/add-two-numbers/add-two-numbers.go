package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func linkedListToNumber(l *ListNode) int {
	mul := 1
	result := 0
	curr := l
	for {
		result += curr.Val * mul
		if curr.Next == nil {
			break
		}
		curr = curr.Next
		mul *= 10
	}
	return result
}

func numberToLinkedList(num int) *ListNode {
	l := &ListNode{}
	curr := l
	for num != 0 {
		r := num % 10
		num /= 10
		curr.Val = r
		if num == 0 {
			break
		}
		curr.Next = &ListNode{}
		curr = curr.Next
	}
	return l
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l := &ListNode{}

	currL1 := l1
	currL2 := l2
	currL := l
	incr := 0

	for {
		if currL1 != nil {
			currL.Val += currL1.Val
			currL1 = currL1.Next
		}
		if currL2 != nil {
			currL.Val += currL2.Val
			currL2 = currL2.Next
		}
		if incr > 0 {
			currL.Val += incr
			incr = 0
		}
		if currL.Val >= 10 {
			incr = currL.Val / 10
			currL.Val %= 10
		}
		if currL1 == nil && currL2 == nil {
			if incr != 0 {
				currL.Next = &ListNode{Val: incr}
			}
			break
		}
		currL.Next = &ListNode{}
		currL = currL.Next
	}
	return l
}

func print(l *ListNode) {
	curr := l
	for {
		fmt.Printf("%d", curr.Val)
		if curr.Next == nil {
			break
		}
		curr = curr.Next
		fmt.Print("->")
	}
	fmt.Println()
}

func main() {
	print(numberToLinkedList(342))
	print(addTwoNumbers(numberToLinkedList(5), numberToLinkedList(5)))
}
