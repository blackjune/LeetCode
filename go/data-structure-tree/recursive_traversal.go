package main

import (
	"fmt"
	"leetcode/common"
)

func preorderTraversal(root *common.TreeNode) []string {
	fmt.Println(root.Val)
	if root.Left != nil {
		preorderTraversal(root.Left)
	}
	if root.Right != nil {
		preorderTraversal(root.Right)
	}
	return []string{}
}

func postorderTraversal(root *common.TreeNode) []string {
	if root.Left != nil {
		postorderTraversal(root.Left)
	}
	if root.Right != nil {
		postorderTraversal(root.Right)
	}
	fmt.Println(root.Val)
	return []string{}
}

func inorderTraversal(root *common.TreeNode) []string {
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
	tree := common.TreeNode{
		Val: "F",
		Left: &common.TreeNode{
			Val: "B",
			Left: &common.TreeNode{
				Val: "A",
			},
			Right: &common.TreeNode{
				Val: "D",
				Left: &common.TreeNode{
					Val: "C",
				},
				Right: &common.TreeNode{
					Val: "E",
				},
			},
		},
		Right: &common.TreeNode{
			Val: "G",
			Right: &common.TreeNode{
				Val: "I",
				Left: &common.TreeNode{
					Val: "H",
				},
			},
		},
	}
	fmt.Println(inorderTraversal(&tree))
}
