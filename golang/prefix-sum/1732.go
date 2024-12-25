package main

func largestAltitude(gain []int) int {
	initAlt, hstAlt := 0, 0

	for _, g := range gain {
		initAlt += g
		if initAlt > hstAlt {
			hstAlt = initAlt
		}
	}

	return hstAlt
}
