package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	} else {
		return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
