package main

import (
	"container/heap"
	"sort"
)

// Maximum Subsequence Score

// IntHeap ...
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push ...
func (h *IntHeap) Push(ele interface{}) { *h = append(*h, ele.(int)) }

// Pop ...
func (h *IntHeap) Pop() interface{} {
	n := len(*h)
	old := *h
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	n := len(nums1)

	pairs := make([][2]int, n)
	for i := 0; i < n; i++ {
		pairs[i] = [2]int{nums1[i], nums2[i]}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] > pairs[j][1]
	})

	h := &IntHeap{}
	heap.Init(h)

	total, resp := 0, 0
	for _, pair := range pairs {
		a, b := pair[0], pair[1]
		heap.Push(h, a)
		total += a

		if h.Len() > k {
			total -= heap.Pop(h).(int)
		}

		if h.Len() == k {
			pdt := total * b
			if pdt > resp {
				resp = pdt
			}
		}
	}

	return int64(resp)
}
