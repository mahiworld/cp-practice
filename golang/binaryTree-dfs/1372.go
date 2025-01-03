package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) int {
	maxLength := 0

	if root.Left != nil {
		dfsTraverse(root.Left, "left", 1, &maxLength)
	}

	if root.Right != nil {
		dfsTraverse(root.Right, "right", 1, &maxLength)
	}

	return maxLength
}

func dfsTraverse(node *TreeNode, dir string, length int, maxLength *int) {
	if node == nil {
		return
	}

	if length > *maxLength {
		*maxLength = length
	}

	if dir == "left" {
		dfsTraverse(node.Right, "right", length+1, maxLength)
		dfsTraverse(node.Left, "left", 1, maxLength)
	}

	if dir == "right" {
		dfsTraverse(node.Right, "right", 1, maxLength)
		dfsTraverse(node.Left, "left", length+1, maxLength)
	}
}
