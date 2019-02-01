package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	node := head
	for node.Next != nil {
		if node.Next.Val == node.Val {
			node.Next = node.Next.Next
		}
		node = node.Next
	}
	return head
}
