package main

//N-th Tribonacci Number

// Approach:1 Recursion ------------------------
// func tribonacci(n int) int {
// 	if n == 0 {
// 		return 0
// 	} else if n == 1 || n == 2 {
// 		return 1
// 	}

// 	return tribonacci(n-3) + tribonacci(n-2) + tribonacci(n-1)
// }

// Apprach:2 Iterative ---------------------------
func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	t0, t1, t2 := 0, 1, 1
	for i := 3; i <= n; i++ {
		t3 := t0 + t1 + t2
		t0, t1, t2 = t1, t2, t3
	}

	return t2
}
