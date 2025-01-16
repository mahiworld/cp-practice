package main

import "math"

// Smallest Number in Infinite Set
type SmallestInfiniteSet struct {
	cur int
	set map[int]bool
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		cur: 1,
		set: make(map[int]bool),
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if len(this.set) > 0 {
		// Find the smallest element in the set
		smallest := math.MaxInt
		for num := range this.set {
			if num < smallest {
				smallest = num
			}
		}

		// Remove the smallest element from the set
		delete(this.set, smallest)

		return smallest
	} else {
		this.cur++
		return this.cur - 1
	}
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if this.cur > num {
		this.set[num] = true
	}
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
