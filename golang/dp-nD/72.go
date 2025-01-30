package main

// Edit Distance
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	// State Definition: Create a 2d DP array
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Base Case: Initialize base cases
	for i := 0; i <= m; i++ {
		dp[i][0] = i // If word2 is empty (j == 0), we need to delete all characters of word1
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j // If word1 is empty (i == 0), we need to insert all characters of word2
	}

	// State Transition: Fill DP table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// If characters do not match, consider all three operations
				dp[i][j] = min(
					dp[i-1][j],   // Delete
					dp[i][j-1],   // Insert
					dp[i-1][j-1], // Replace
				) + 1
			}
		}
	}

	// Result Extraction:: Return the result in dp[m][n]
	return dp[m][n]
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}

	if b < c {
		return b
	}
	return c
}
