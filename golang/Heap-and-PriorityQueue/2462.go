package main

import "container/heap"

//Total Cost to Hire K Workers

// func totalCost(costs []int, k int, candidates int) int64 {
//     resp := 0
//     slice := append([]int{}, costs...)

//     for k > 0 {
//         minCost := math.MaxInt
//         minIndex := -1

//         n := len(slice)
//         left, right := 0, n-1

//         for left < candidates && left <= right {
//             if slice[left] < minCost {
//                 minCost = slice[left]
//                 minIndex = left
//             }
//             left++
//         }

//         for right >= n-candidates && right >= left {
//             if slice[right] < minCost {
//                 minCost = slice[right]
//                 minIndex = right
//             }
//             right--
//         }

//         resp += minCost
//         slice = append(slice[:minIndex], slice[minIndex+1:]...)
//         k--
//     }

//     return int64(resp)
// }

type Worker struct {
	cost  int
	index int
}

type MinHeap []Worker

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Worker))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func totalCost(costs []int, k int, candidates int) int64 {
	leftHeap := &MinHeap{}
	rightHeap := &MinHeap{}
	n := len(costs)

	for i := 0; i < candidates && i < n; i++ {
		heap.Push(leftHeap, Worker{costs[i], i})
	}
	for i := n - 1; i >= n-candidates && i >= candidates; i-- {
		heap.Push(rightHeap, Worker{costs[i], i})
	}

	hired := 0
	resp := int64(0)
	leftPointer, rightPointer := candidates, n-candidates-1

	for hired < k {
		var selected Worker

		if leftHeap.Len() > 0 && (rightHeap.Len() == 0 || (*leftHeap)[0].cost <= (*rightHeap)[0].cost) {
			selected = heap.Pop(leftHeap).(Worker)
			if leftPointer <= rightPointer {
				heap.Push(leftHeap, Worker{costs[leftPointer], leftPointer})
				leftPointer++
			}
		} else if rightHeap.Len() > 0 {
			selected = heap.Pop(rightHeap).(Worker)
			if rightPointer >= leftPointer {
				heap.Push(rightHeap, Worker{costs[rightPointer], rightPointer})
				rightPointer--
			}
		}

		resp += int64(selected.cost)
		hired++
	}

	return resp
}
