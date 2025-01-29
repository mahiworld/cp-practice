package main

// Best Time to Buy and Sell Stock with Transaction Fee

// Approach : 1 -------------------------------------------------
// func maxProfit(prices []int, fee int) int {
// 	n := len(prices)

// 	// State Definition: Create a 2D dp array
// 	dp := make([][]int, n+1)
// 	for i := range dp {
// 		dp[i] = make([]int, 2) // 0: not holding, 1: holding
// 	}

// 	// State Transition: Fill the DP array
// 	dp[0][0], dp[0][1] = 0, -prices[0]
// 	for i := 1; i < n; i++ {
// 		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee) // Sell
// 		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])     // Buy
// 	}

// 	// Result Extraction: Maximum profit without holding stock at the end
// 	return dp[n-1][0]
// }

// Approach : 2 ---------------------------------------------------
func maxProfit(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	hold, cash := -prices[0], 0

	for i := 1; i < n; i++ {
		cash = max(cash, hold+prices[i]-fee) // Update cash (profit when not holding stock) by selling
		hold = max(hold, cash-prices[i])     // Update hold (profit when holding stock) by buying
	}

	return cash
}
