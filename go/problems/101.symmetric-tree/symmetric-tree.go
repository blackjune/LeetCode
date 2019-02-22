package tree

import (
	"leetcode/common"
)

func isSymmetric(root *common.TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left == nil && root.Right != nil {
		return false
	}
	if root.Left != nil && root.Right == nil {
		return false
	}
	if root.Left.Val != root.Right.Val {
		return false
	}
	return isSymmetric(root.Left) && isSymmetric(root.Right)
}
