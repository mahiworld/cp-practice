package main

func longestSubarray(nums []int) int {
	left := 0
	zeros := 1
	result := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeros--
		}

		for zeros < 0 {
			if nums[left] == 0 {
				zeros++
			}
			left++
		}

		result = max(result, right-left+1)
	}

	return result - 1
}
