package main

//N-th Tribonacci Number

// Approach:1 Recursion ------------------------
func tribonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}

	return tribonacci(n-3) + tribonacci(n-2) + tribonacci(n-1)
}
