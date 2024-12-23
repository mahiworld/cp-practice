package main

func uniqueOccurrences(arr []int) bool {
	ocrMap := make(map[int]int)
	checkMap := make(map[int]bool)

	for _, ele := range arr {
		ocrMap[ele]++
	}

	for _, val := range ocrMap {
		if checkMap[val] {
			return false
		}
		checkMap[val] = true
	}

	return true
}
