package main

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1

	result := 0

	for i < j {
		result = max(result, (j-i)*min(height[i], height[j]))

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return result
}

func max(m, n int) int {
	if m > n {
		return m
	}

	return n
}

func min(m, n int) int {
	if m < n {
		return m
	}

	return n
}
