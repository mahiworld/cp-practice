package main

import "sort"

// func closeStrings(word1 string, word2 string) bool {
//     if len(word1) != len(word2) {
//         return false
//     }

//     // Frequency and character presence maps
//     w1Freq, w2Freq := make(map[rune]int), make(map[rune]int)
//     w1Set, w2Set := make(map[rune]bool), make(map[rune]bool)

//     for _, char := range word1 {
//         w1Freq[char]++
//         w1Set[char] = true
//     }

//     for _, char := range word2 {
//         w2Freq[char]++
//         w2Set[char] = true
//     }

//     // Check if character sets are equal
//     if len(w1Set) != len(w2Set) {
//         return false
//     }

//     for char := range w1Set {
//         if !w2Set[char] {
//             return false
//         }
//     }

//     // Collect frequency values and sort them
//     freq1, freq2 := make([]int, 0, len(w1Freq)), make([]int, 0, len(w2Freq))
//     for _, count := range w1Freq {
//         freq1 = append(freq1, count)
//     }
//     for _, count := range w2Freq {
//         freq2 = append(freq2, count)
//     }

//     sort.Ints(freq1)
//     sort.Ints(freq2)

//     // Compare sorted frequency lists
//     for i := range freq1 {
//         if freq1[i] != freq2[i] {
//             return false
//         }
//     }

//     return true
// }

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	// Arrays for frequency counts (assuming lowercase English letters)
	freq1, freq2 := make([]int, 26), make([]int, 26)
	chars1, chars2 := make([]bool, 26), make([]bool, 26)

	for _, char := range word1 {
		index := char - 'a'
		freq1[index]++
		chars1[index] = true
	}

	for _, char := range word2 {
		index := char - 'a'
		freq2[index]++
		chars2[index] = true
	}

	// Check if character sets are the same
	for i := 0; i < 26; i++ {
		if chars1[i] != chars2[i] {
			return false
		}
	}

	// Sort frequencies and compare
	sort.Ints(freq1)
	sort.Ints(freq2)

	for i := 0; i < 26; i++ {
		if freq1[i] != freq2[i] {
			return false
		}
	}

	return true
}
