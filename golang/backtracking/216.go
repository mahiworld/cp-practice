package main

// Combination Sum III
func combinationSum3(k int, n int) [][]int {
	var result [][]int

	backtrack(&result, []int{}, 1, k, n, 0)

	return result
}

func backtrack(result *[][]int, curr []int, start, k, n, sum int) {
	if len(curr) == k && sum == n {
		temp := make([]int, len(curr))
		copy(temp, curr)
		*result = append(*result, temp)
		return
	}

	for i := start; i <= 9 && sum+i <= n; i++ {
		curr = append(curr, i)
		backtrack(result, curr, i+1, k, n, sum+i)
		curr = curr[:len(curr)-1]
	}
}
