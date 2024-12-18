package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := candies[0]
	for _, cand := range candies {
		if cand > max {
			max = cand
		}
	}

	result := make([]bool, 0)
	for _, cand := range candies {
		if cand+extraCandies >= max {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}

	return result
}
