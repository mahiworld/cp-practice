package main

import "strconv"

func compress(chars []byte) int {
	n := len(chars)
	i, j := 0, 0

	for j < n {
		char := chars[j]
		count := 0

		for j < n && chars[j] == char {
			j++
			count++
		}

		chars[i] = char
		i++

		if count > 1 {
			countStr := strconv.Itoa(count)
			for c := 0; c < len(countStr); c++ {
				chars[i] = countStr[c]
				i++
			}
		}
	}

	return i
}
