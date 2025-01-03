package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
	prefixSums := map[int]int{0: 1}
	return dfs(root, 0, targetSum, prefixSums)
}

func dfs(node *TreeNode, csum, tsum int, psums map[int]int) int {
	if node == nil {
		return 0
	}

	csum += node.Val

	diff := csum - tsum
	count := psums[diff]

	psums[csum]++

	count += dfs(node.Left, csum, tsum, psums)
	count += dfs(node.Right, csum, tsum, psums)

	//on backtrack remove from psums
	psums[csum]--
	if psums[csum] == 0 {
		delete(psums, csum)
	}

	return count
}
