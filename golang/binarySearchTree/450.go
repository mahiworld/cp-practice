package main

// Delete Node in a BST
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key == root.Val {
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Left == nil || root.Right == nil {
			child := root.Left
			if root.Right != nil {
				child = root.Right
			}
			return child
		} else {
			successor := findMin(root.Right)
			root.Val = successor.Val
			root.Right = deleteNode(root.Right, successor.Val)
		}
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}

	return root
}

func findMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}
