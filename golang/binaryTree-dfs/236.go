package main

// Longest common ancestor of a binary tree

// Approach - 2
// 1. The function lowestCommonAncestor takes in three parameters: the root of a binary tree (root) and two nodes of the binary tree (p and q).
// 2. The first if statement checks if the root is None or if it is equal to either p or q. If either of these conditions is true, it means that
// we have found one of the nodes we are looking for, and we return the root.
// 3. Next, we recursively call the lowestCommonAncestor function on the left and right subtrees of the root, passing in the same nodes p and q.
// We store the results of these recursive calls in variables l and r, respectively.
// 4. The second if statement checks if both l and r are not None. If this condition is true, it means that we have found both p and q in different
// subtrees of the current root, and therefore the current root is the lowest common ancestor. We return the current root.
// 5. If the second if statement is not satisfied, we return either l or r, depending on which one is not None. This is because if only one of l and r is not None,
// it means that the other node is not in the subtree of the current root, so we return the node that is in the subtree.
// 6. If none of the previous conditions is satisfied, it means that both l and r are None, so we return None. This happens when we have reached the
// end of a branch of the binary tree without finding either p or q.
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	return right
}

//--------------------------------------
//Approach - 1
// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
//     // Map to store ancestors of all nodes
//     ansts := make(map[int][]*TreeNode)

//     // Collect ancestors of all nodes
//     findAncestors(root, p, q, nil, &ansts)

//     // If ancestors of p or q are missing, return nil
//     if _, foundP := ansts[p.Val]; !foundP {
//         return nil
//     }
//     if _, foundQ := ansts[q.Val]; !foundQ {
//         return nil
//     }

//     // Get the ancestors of p and q
//     anstsOfP := ansts[p.Val]
//     anstsOfQ := ansts[q.Val]

//     //-----This is not correct, it assumes that the smallest-valued common ancestor
//     //-----(by Val) is always the correct LCA, which is not true.
//     //-----The correct LCA is the lowest (deepest) node in the tree that is a common ancestor
//     //-----of both
//     // // Use a map for faster lookups
//     // ancestorsSet := make(map[int]bool)
//     // for _, ancestor := range anstsOfP {
//     //     ancestorsSet[ancestor.Val] = true
//     // }

//     // // Find the common ancestor with the smallest value
//     // var lca *TreeNode
//     // minVal := math.MaxInt
//     // for _, ancestor := range anstsOfQ {
//     //     if ancestorsSet[ancestor.Val] && ancestor.Val < minVal {
//     //         lca = ancestor
//     //         minVal = ancestor.Val
//     //     }
//     // }

//     // return lca
//     // Find the lowest common ancestor
//     lca := findLowestCommonAncestor(anstsOfP, anstsOfQ)
//     return lca
// }

// func findAncestors(node, p, q *TreeNode, path []*TreeNode, ansts *map[int][]*TreeNode) {
//     if node == nil {
//         return
//     }

//     // Add the current node to the path
//     path = append(path, node)

//     // Store the path (ancestors) for the current node
//     (*ansts)[node.Val] = append([]*TreeNode{}, path...)

//     // Stop if both p and q have ancestors
//     if _, foundP := (*ansts)[p.Val]; foundP {
//         if _, foundQ := (*ansts)[q.Val]; foundQ {
//             return
//         }
//     }

//     // Continue traversing the tree
//     findAncestors(node.Left, p, q, path, ansts)
//     findAncestors(node.Right, p, q, path, ansts)
// }

// func findLowestCommonAncestor(anstsOfP, anstsOfQ []*TreeNode) *TreeNode {
//     minLen := len(anstsOfP)
//     if len(anstsOfQ) < minLen {
//         minLen = len(anstsOfQ)
//     }

//     var lca *TreeNode
//     for i := 0; i < minLen; i++ {
//         if anstsOfP[i] == anstsOfQ[i] {
//             lca = anstsOfP[i]
//         } else {
//             break
//         }
//     }
//     return lca
// }
