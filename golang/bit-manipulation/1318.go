package main

// Minimum Flips to Make a OR b Equal to c
func minFlips(a int, b int, c int) int {
	flips := 0
	for a > 0 || b > 0 || c > 0 {
		bitA := a & 1
		bitB := b & 1
		bitC := c & 1

		if (bitA | bitB) != bitC {
			if bitC == 1 {
				flips++
			} else {
				flips += bitA + bitB
			}
		}

		a >>= 1
		b >>= 1
		c >>= 1
	}
	return flips
}
