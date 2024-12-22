package main

func productExceptSelf(nums []int) []int {

	n := len(nums)

	result := make([]int, n)

	prefix, suffix := 1, 1

	//prefix pdt
	for i := 0; i < n; i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	//suffix pdt
	for i := n - 1; i >= 0; i-- {
		result[i] *= suffix // multiply suffix with prefix
		suffix *= nums[i]
	}

	return result
}
