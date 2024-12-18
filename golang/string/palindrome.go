package main

func isPalindrome(s string) bool {
	runes := []rune(s)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if runes[1] != runes[j] {
			return false
		}
	}

	return true
}
