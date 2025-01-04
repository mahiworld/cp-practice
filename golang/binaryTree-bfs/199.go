package main

//Binary Tree Right Side View
// func rightSideView(root *TreeNode) []int {
//     if root == nil{
//         return nil
//     }

// 	type NodeLevel struct {
// 		Node  *TreeNode
// 		Level int
// 	}

// 	queue := []NodeLevel{{Node: root, Level: 0}}
//     resMap := map[int]int{0: root.Val}

//     for len(queue) > 0{
//         levelnode := queue[0]
//         node := levelnode.Node
//         queue = queue[1:]
//         level := levelnode.Level

//         flag := false
//         if node.Left != nil{
//             queue = append(queue, NodeLevel{Node: node.Left, Level: level+1})
//         }

//         if node.Right != nil{
//             flag = true
//             queue = append(queue, NodeLevel{Node: node.Right, Level: level+1})
//         }

//         if flag{
//             resMap[level+1] = node.Right.Val
//         }else{
//             if node.Left != nil{
//              resMap[level+1] = node.Left.Val
//             }
//         }
//     }

//     res := []int{}
// 	for i := 0; i < len(resMap); i++ {
// 		if val, ok := resMap[i]; ok {
// 			res = append(res, val)
// 		}
// 	}

//     return res
// }

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	res := []int{}

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if i == levelSize-1 {
				res = append(res, node.Val)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return res
}
