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
