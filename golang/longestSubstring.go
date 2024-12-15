package main

// longest substring without repeating characters ----- sliding window problem
func longestSubstring(s string) int {
	charIndex := make(map[byte]int) // make(map[rune]int)
	start, result := 0, 0

	// for index, ch := range s { // here ch will be rune type
	// 	if foundIndex, ok := charIndex[ch]; ok && foundIndex >= start {
	// 		start = foundIndex + 1
	// 	}

	// 	charIndex[ch] = index
	// result = max(result, index-start+1)
	// }

	for end := 0; end < len(s); end++ {
		if foundIndex, ok := charIndex[s[end]]; ok && foundIndex >= start { //foundIndex >= start explicitly ensures that we adjust start only when the duplicate character appears inside the window, avoiding unnecessary adjustments that would break the sliding window logic.
			start = foundIndex + 1
		}

		charIndex[s[end]] = end
		result = max(result, end-start+1)
	}

	return result
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
