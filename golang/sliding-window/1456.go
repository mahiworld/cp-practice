package main

func maxVowels(s string, k int) int {
	max, j := 0, 0

	result := max

	for j < len(s) {
		if isVowel(s[j]) {
			max++
		}

		if j >= k {
			if isVowel(s[j-k]) {
				max--
			}
		}

		if max > result {
			result = max
		}

		j++
	}

	return result
}

func isVowel(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' ||
		ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U'
}
