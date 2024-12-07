package main

import (
	"sort"
	"strings"
)

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func wordFrequency(s string) map[string]int {
	words := strings.Fields(s)
	freq := make(map[string]int)

	for _, word := range words {
		freq[word]++
	}

	return freq
}

func sortString(s string) string {
	runes := []rune(s)

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return string(runes)
}
