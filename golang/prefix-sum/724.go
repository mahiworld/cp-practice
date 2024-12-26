package main

func pivotIndex(nums []int) int {
	leftSum := 0
	rightSum := 0
	for _, num := range nums {
		rightSum += num
	}

	for i := 0; i < len(nums); i++ {
		rightSum -= nums[i]

		if leftSum == rightSum {
			return i
		}

		leftSum += nums[i]
	}

	return -1
}
