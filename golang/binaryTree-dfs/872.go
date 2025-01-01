package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leafs1 := []int{}
	leafs2 := []int{}

	findLeafs(root1, &leafs1)
	findLeafs(root2, &leafs2)

	if len(leafs1) != len(leafs2) {
		return false
	}
	for index, ele := range leafs1 {
		if ele != leafs2[index] {
			return false
		}
	}
	return true
}

func findLeafs(root *TreeNode, leafs *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*leafs = append(*leafs, root.Val)
	} else {
		findLeafs(root.Left, leafs)
		findLeafs(root.Right, leafs)
	}
}

// func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
//     ch1 := make(chan []int)
//     ch2 := make(chan []int)

//     go calculateLeafs(root1, ch1)
//     go calculateLeafs(root2, ch2)

//     leafs1 := <-ch1
//     leafs2 := <-ch2

//     if len(leafs1) != len(leafs2) {
//         return false
//     }
//     for index, ele := range leafs1 {
//         if ele != leafs2[index] {
//             return false
//         }
//     }
//     return true
// }

// func calculateLeafs(root *TreeNode, ch chan []int){
//     leafs := []int{}
//     findLeafs(root, &leafs)
//     ch <- leafs
// }
