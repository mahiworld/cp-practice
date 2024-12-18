package main

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ { //i * i <= n instead of i <= n is because a larger factor of n must be paired with a smaller factor that has already been checked
		if n%i == 0 {
			return false
		}
	}
	return true
}
