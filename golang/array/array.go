package main

func mergeSortedArrays(arr1, arr2 []int) []int {
	arr := make([]int, 0, len(arr1)+len(arr2))

	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			arr = append(arr, arr1[i])
			i++
		} else {
			arr = append(arr, arr2[j])
			j++
		}
	}
	arr = append(arr, arr1[i:]...)
	arr = append(arr, arr2[j:]...)

	return arr
}

func isAscending(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

// rearrange as [max, min, second_max, second_min, ...] pattern ------ 2 pointers problem
func rearrangeOutPlace(arr []int) []int {
	newArr := make([]int, 0)
	n := len(arr)

	if isAscending(arr) {
		for i := 0; i < n/2; i++ {
			newArr = append(newArr, arr[n-1-i], arr[i])
		}
	} else {
		for i := 0; i < n/2; i++ {
			newArr = append(newArr, arr[i], arr[n-1-i])
		}
	}

	// If the length is odd, append the middle element
	if n%2 != 0 {
		newArr = append(newArr, arr[n/2])
	}

	return newArr
}
