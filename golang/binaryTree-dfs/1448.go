package main

import "fmt"

func insertLevelOrder(arr []int, root *TreeNode, i int) *TreeNode {
	if i < len(arr) {
		temp := &TreeNode{Val: arr[i]}
		root = temp

		// Insert left child
		root.Left = insertLevelOrder(arr, root.Left, 2*i+1)

		// Insert right child
		root.Right = insertLevelOrder(arr, root.Right, 2*i+2)
	}
	return root
}

func main() {
	arr := []int{3, 1, 4, 3, 0, 1, 5}

	var root *TreeNode
	root = insertLevelOrder(arr, root, 0)

	fmt.Println(goodNodes(root))
}

// Approach:
// Recursively DFT traverses the tree
// counting nodes that are greater than or equal to the maximum value encountered along the path.
func goodNodes(root *TreeNode) int {
	return calGoodNodes(root, root.Val)
}

func calGoodNodes(node *TreeNode, max int) int {
	if node == nil {
		return 0
	}

	count := 0
	if node.Val >= max {
		count = 1
		max = node.Val
	}

	count += calGoodNodes(node.Left, max)
	count += calGoodNodes(node.Right, max)

	return count
}
