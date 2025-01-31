package main

import (
	"strconv"
	"strings"
)

// Counting Bits

// Approach : 1 -------------------------------
func countBits(n int) []int {
	resp := make([]int, 0)

	for i := 0; i <= n; i++ {
		resp = append(resp, countBinaryOnes(i))
	}

	return resp
}

func countBinaryOnes(num int) int {
	binaryStr := strconv.FormatInt(int64(num), 2)

	return strings.Count(binaryStr, "1")
}
