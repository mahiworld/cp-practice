package main

//Longest Common Subsequence

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)

	// State Definition: Create a 2D dp array
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// State Transition: Fill the DP array
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] { // Compare characters from text1 and text2
				dp[i][j] = dp[i-1][j-1] + 1 // Include the character
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1]) // Exclude one character
			}
		}
	}

	// Result Extraction: The bottom-right corner contains the length of LCS
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
