package main

// Counting Bits

// Approach : 1 -------------------------------
// func countBits(n int) []int {
// 	resp := make([]int, 0)

// 	for i := 0; i <= n; i++ {
// 		resp = append(resp, countBinaryOnes(i))
// 	}

// 	return resp
// }

// func countBinaryOnes(num int) int {
// 	binaryStr := strconv.FormatInt(int64(num), 2)

// 	return strings.Count(binaryStr, "1")
// }

// Approach : 2 -> Dynamic programming approach
func countBits(n int) []int {
	resp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		// i >> 1 shifts the bits right, effectively removing the least significant bit
		// i & 1 checks whether the least significant bit is 1
		resp[i] = resp[i>>1] + (i & 1)
	}

	return resp
}
