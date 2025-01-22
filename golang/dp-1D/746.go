package main

// Min Cost Climbing Stairs
func minCostClimbingStairs(cost []int) int {
	n := len(cost)

	dp := make([]int, n+1)

	dp[n] = 0

	for i := n - 1; i >= 0; i-- {
		if i+2 <= n {
			dp[i] = cost[i] + min(dp[i+1], dp[i+2])
		} else {
			dp[i] = cost[i] + dp[i+1]
		}
	}

	return min(dp[0], dp[1])
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
