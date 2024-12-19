package main

import "strings"

func reverseVowels(s string) string {
	runes := []rune(s)

	i, j := 0, len(s)-1
	for i < j {

		//find vowel from left
		if !isVowel(runes[i]) {
			i++
			continue
		}

		//find vowel from right
		if !isVowel(runes[j]) {
			j--
			continue
		}

		//swap the vowels
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}

	return string(runes)
}

func isVowel(r rune) bool {
	//  return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
	//        c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
	return strings.ContainsRune("aeiouAEIOU", r)
}
