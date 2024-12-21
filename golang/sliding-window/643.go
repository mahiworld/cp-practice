package main

// sliding window technique
func findMaxAverage(nums []int, k int) float64 {
	if len(nums) < k {
		return 0
	}

	maxTotal := total(nums[:k])
	newMax := maxTotal

	for i := 0; i < len(nums)-k; i++ {

		newMax -= nums[i]
		newMax += nums[i+k]

		maxTotal = max(maxTotal, newMax)
	}

	return float64(maxTotal) / float64(k)
}

func total(nums []int) int {
	total := 0
	for _, ele := range nums {
		total += ele
	}

	return total
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

// func findMaxAverage(nums []int, k int) float64 {
//     // Calculate the initial sum of the first `k` elements
//     currentSum := 0
//     for i := 0; i < k; i++ {
//         currentSum += nums[i]
//     }

//     maxSum := currentSum

//     // Slide the window over the array
//     for i := k; i < len(nums); i++ {
//         currentSum = currentSum - nums[i-k] + nums[i]
//         if currentSum > maxSum {
//             maxSum = currentSum
//         }
//     }

//     // Calculate the maximum average
//     return float64(maxSum) / float64(k)
// }
