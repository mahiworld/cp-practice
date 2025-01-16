package main

import "math"

// Smallest Number in Infinite Set

// SmallestInfiniteSet ...
type SmallestInfiniteSet struct {
	cur int
	set map[int]bool
}

// Constructor ...
func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		cur: 1,
		set: make(map[int]bool),
	}
}

// PopSmallest ...
func (s *SmallestInfiniteSet) PopSmallest() int {
	if len(s.set) > 0 {
		// Find the smallest element in the set
		smallest := math.MaxInt
		for num := range s.set {
			if num < smallest {
				smallest = num
			}
		}

		// Remove the smallest element from the set
		delete(s.set, smallest)

		return smallest
	}

	s.cur++
	return s.cur - 1
}

// AddBack ...
func (s *SmallestInfiniteSet) AddBack(num int) {
	if s.cur > num {
		s.set[num] = true
	}
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
