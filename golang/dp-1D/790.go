package main

// Domino and Tromino Tiling
func numTilings(n int) int {
	const mod = 1_000_000_007

	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 5
	}

	dp := make([]int, n+1)

	dp[1] = 1
	dp[2] = 2
	dp[3] = 5

	for i := 4; i <= n; i++ {
		dp[i] = (2*dp[i-1] + dp[i-3]) % mod
	}

	return dp[n]
}
