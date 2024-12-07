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
