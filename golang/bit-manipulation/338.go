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

// Approach : 2 ----------------------------------
func countBits(n int) []int {
	resp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		resp[i] = resp[i>>1] + (i & 1)
	}

	return resp
}
