package main

//Maximum Level Sum of a Binary Tree

// func maxLevelSum(root *TreeNode) int {
// 	type NodeLevel struct {
// 		Node  *TreeNode
// 		Level int
// 	}

// 	type Result struct {
// 		Level int
// 		Sum   int
// 	}

//     var res *Result

// 	queue := []NodeLevel{{Node: root, Level: 1}}

// 	for len(queue) > 0 {
// 		levelSize := len(queue)
// 		levelSum := 0
// 		level := 1
// 		for i := 0; i < levelSize; i++ {
// 			node := queue[0].Node
// 			level = queue[0].Level
// 			queue = queue[1:]
// 			levelSum += node.Val

// 			if node.Left != nil {
// 				queue = append(queue, NodeLevel{Node: node.Left, Level: level + 1})
// 			}

// 			if node.Right != nil {
// 				queue = append(queue, NodeLevel{Node: node.Right, Level: level + 1})
// 			}

// 		}

// 		if res == nil {
//             res = &Result{
//                 Level: level,
//                 Sum: levelSum,
//             }
// 		} else {
// 			if levelSum > res.Sum {
// 				res.Level = level
// 				res.Sum = levelSum
// 			}
// 		}
// 	}

// 	return res.Level
// }

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	level := 1
	maxLevel := 1
	maxSum := root.Val

	for len(queue) > 0 {
		levelSize := len(queue)
		levelSum := 0

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			levelSum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if levelSum > maxSum {
			maxSum = levelSum
			maxLevel = level
		}

		level++
	}

	return maxLevel
}
