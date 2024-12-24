package main

import "math"

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	// first, second := int(^uint(0)>>1), int(^uint(0)>>1)
	first, second := math.MaxInt, math.MaxInt

	for _, num := range nums {
		if num <= first {
			first = num
		} else if num <= second {
			second = num
		} else {
			return true
		}
	}

	return false
}
