package main

// House Robber
// Apprach 1 -------------------------
func rob(nums []int) int {
	var prev1, prev2 int
	for _, num := range nums {
		current := max(prev1, prev2+num)
		prev2 = prev1
		prev1 = current
	}

	return prev1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

//Apprach:2 -------------
// func rob(nums []int) int {
//     n := len(nums)

//     if n == 0 {
//         return 0
//     }
//     if n == 1 {
//         return nums[0]
//     }

//     dp := make([]int, n)

//     dp[0] = nums[0]
//     dp[1] = max(nums[0], nums[1])

//     for i := 2; i<n; i++{
//         dp[i] = max(dp[i-1], nums[i]+dp[i-2])
//     }

//     return dp[n-1]
// }
