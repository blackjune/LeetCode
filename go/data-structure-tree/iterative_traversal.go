package main

import (
	"fmt"
)

type TreeNode struct {
	Val   string
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []string {
	curr := root
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Left
	}
	return []string{}
}

func postorderTraversal(root *TreeNode) []string {
	if root.Left != nil {
		postorderTraversal(root.Left)
	}
	if root.Right != nil {
		postorderTraversal(root.Right)
	}
	fmt.Println(root.Val)
	return []string{}
}

func inorderTraversal(root *TreeNode) []string {
	if root.Left != nil {
		inorderTraversal(root.Left)
	}
	fmt.Println(root.Val)
	if root.Right != nil {
		inorderTraversal(root.Right)
	}
	return []string{}
}

func main() {
	tree := TreeNode{
		Val: "F",
		Left: &TreeNode{
			Val: "B",
			Left: &TreeNode{
				Val: "A",
			},
			Right: &TreeNode{
				Val: "D",
				Left: &TreeNode{
					Val: "C",
				},
				Right: &TreeNode{
					Val: "E",
				},
			},
		},
		Right: &TreeNode{
			Val: "G",
			Right: &TreeNode{
				Val: "I",
				Left: &TreeNode{
					Val: "H",
				},
			},
		},
	}
	fmt.Println(preorderTraversal(&tree))
}
