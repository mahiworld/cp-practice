package main

// Single Number
// func singleNumber(nums []int) int {
// 	visited := make(map[int]bool)
// 	for _, num := range nums {
// 		if visited[num] {
// 			visited[num] = false
// 		} else {
// 			visited[num] = true
// 		}
// 	}

// 	for key, value := range visited {
// 		if value {
// 			return key
// 		}
// 	}

//		return -1
//	}

func singleNumber(nums []int) int {
	var resp int

	for _, num := range nums {
		resp ^= num
	}

	return resp
}
