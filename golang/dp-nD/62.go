package main

//Unique Paths

func uniquePaths(m int, n int) int {
	//State Definition: create 2d dp array
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	//Base Case: Initialize the base cases
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	//State Transition: Fill the DP array
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	//Result Extraction:
	return dp[m-1][n-1]
}
