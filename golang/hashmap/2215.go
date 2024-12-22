package main

//Brute force
// func findDifference(nums1 []int, nums2 []int) [][]int {
//     result := make([][]int, 2)

//     list1 := make([]int, 0)
//     list1Map := make(map[int]bool)
//     for _, num := range nums1{
//         if list1Map[num]{
//             continue
//         }

//         if !isArrayEle(nums2, num){
//             list1 = append(list1, num)
//         }

//         list1Map[num] = true
//     }
//     result[0] = list1

//     list2 := make([]int, 0)
//     list2Map := make(map[int]bool)
//     for _, num := range nums2{
//         if list2Map[num]{
//             continue
//         }

//         if !isArrayEle(nums1, num){
//             list2 = append(list2, num)
//         }

//         list2Map[num] = true
//     }
//     result[1] = list2

//     return result
// }

// func isArrayEle(arr []int, ele int)bool{
//     for _, a := range arr{
//         if ele == a{
//             return true
//         }
//     }

//     return false
// }

// Hash map
func findDifference(nums1 []int, nums2 []int) [][]int {

	m1, m2 := make(map[int]bool), make(map[int]bool)

	for _, n := range nums1 {
		m1[n] = true
	}
	for _, n := range nums2 {
		m2[n] = true
	}

	res := make([][]int, 2)

	for k := range m1 {
		if _, ok := m2[k]; !ok {
			res[0] = append(res[0], k)
		}
	}
	for k := range m2 {
		if _, ok := m1[k]; !ok {
			res[1] = append(res[1], k)
		}
	}
	return res
}
