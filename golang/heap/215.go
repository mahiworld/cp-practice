package main

import "container/heap"

//Kth Largest Element in an Array
//Aproach:1 -----------------------------
// func findKthLargest(nums []int, k int) int {
// 	n := len(nums)
// 	return quickSelect(nums, 0, n-1, n-k)
// }

// func quickSelect(nums []int, l, r, k int) int {
// 	if l <= r {
// 		pivotIndex := partition(nums, l, r)

// 		if pivotIndex == k {
// 			return nums[pivotIndex]
// 		} else if pivotIndex < k {
// 			return quickSelect(nums, pivotIndex+1, r, k)
// 		} else {
// 			return quickSelect(nums, l, pivotIndex-1, k)
// 		}
// 	}
// 	return -1
// }

// func partition(arr []int, p, q int) int {
// 	pivot := arr[p]
// 	i := p
// 	for j := p + 1; j <= q; j++ {
// 		if arr[j] <= pivot {
// 			i++
// 			arr[i], arr[j] = arr[j], arr[i]
// 		}
// 	}
// 	arr[p], arr[i] = arr[i], arr[p]
// 	return i
// }

//Approach:1 Modified -------------------------------
// func findKthLargest(nums []int, k int) int {
// 	return quickselect(nums, 0, len(nums)-1, len(nums)-k)
// }

// func quickselect(nums []int, left, right, k int) int {
// 	if left == right {
// 		return nums[left]
// 	}

// 	// Choose a random pivot index
// 	pivotIndex := left + rand.Intn(right-left+1)
// 	pivotIndex = partition(nums, left, right, pivotIndex)

// 	// Check the position of the pivot
// 	if pivotIndex == k {
// 		return nums[pivotIndex]
// 	} else if pivotIndex < k {
// 		return quickselect(nums, pivotIndex+1, right, k)
// 	} else {
// 		return quickselect(nums, left, pivotIndex-1, k)
// 	}
// }

// func partition(nums []int, left, right, pivotIndex int) int {
// 	pivot := nums[pivotIndex]
// 	// Move pivot to the end
// 	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]
// 	storeIndex := left

// 	// Move all elements smaller than pivot to the left
// 	for i := left; i < right; i++ {
// 		if nums[i] < pivot {
// 			nums[storeIndex], nums[i] = nums[i], nums[storeIndex]
// 			storeIndex++
// 		}
// 	}

// 	// Move pivot to its final place
// 	nums[storeIndex], nums[right] = nums[right], nums[storeIndex]
// 	return storeIndex
// }

// Approach:2 --------------------------------
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	for _, num := range nums {
		heap.Push(minHeap, num)
		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}

	return (*minHeap)[0]
}
