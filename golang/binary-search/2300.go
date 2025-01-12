package main

import "sort"

//Successful Pairs of Spells and Potions
//Approach: Linear search
// func successfulPairs(spells []int, potions []int, success int64) []int {
// 	res := []int{}
// 	for _, spl := range spells {
// 		count := 0
// 		for _, ptn := range potions {
// 			if spl*ptn >= int(success) {
// 				count++
// 			}
// 		}
// 		res = append(res, count)
// 	}

// 	return res
// }

// //Approach 2: Binray search
func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Slice(potions, func(i, j int) bool {
		return potions[i] < potions[j]
	})

	res := make([]int, len(spells))

	for idx, spell := range spells {
		i := binarySearch(potions, spell, int(success))
		if i == -1 {
			res[idx] = 0
		} else {
			res[idx] = len(potions) - i
		}
	}

	return res
}

func binarySearch(potions []int, spell, success int) int {
	left, right := 0, len(potions)-1

	for left <= right {
		mid := left + (right-left)/2
		if spell*potions[mid] >= success {
			if mid == 0 || spell*potions[mid-1] < success {
				return mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
