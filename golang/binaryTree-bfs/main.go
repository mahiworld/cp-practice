package main

import "fmt"

// TreeNode  defines tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertLevelOrder(arr []int, root *TreeNode, i int) *TreeNode {
	if i >= len(arr) {
		return nil
	}

	root = &TreeNode{Val: arr[i]}

	root.Left = insertLevelOrder(arr, root.Left, 2*i+1)
	root.Right = insertLevelOrder(arr, root.Right, 2*i+2)

	return root
}

func main() {
	arr := []int{3, 1, 4, 3, 0, 1, 5}

	var root *TreeNode
	root = insertLevelOrder(arr, root, 0)

	fmt.Println(root)
}
