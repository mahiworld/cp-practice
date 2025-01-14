package main

// Koko Eating Bananas
// Approach : Find kmin & kmax, return in h == len(piles), apply binary search and returm kmin
func minEatingSpeed(piles []int, h int) int {
	kmin, kmax := 1, 1
	for _, ele := range piles {
		if ele > kmax {
			kmax = ele
		}
	}

	if len(piles) == h {
		return kmax
	}

	for kmin < kmax {
		mid := (kmin + kmax) / 2

		totalHours := 0
		for _, ele := range piles {
			totalHours += (ele + mid - 1) / mid
		}
		if totalHours <= h {
			kmax = mid
		} else {
			kmin = mid + 1
		}
	}

	return kmin
}
