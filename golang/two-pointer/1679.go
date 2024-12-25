package main

//Brute force: This is o(n2)
// func maxOperations(nums []int, k int) int {
//     n := len(nums)

//     if n < 2{
//         return 0
//     }

//     count := 0

//     for i:=0; i<n-1; i++{
//         for j:=i+1; j<n;j++{
//             if nums[j] == k-nums[i]{
//                 count++
//                 nums[i] = k+1
//                 nums[j] = k+1
//             }
//         }
//     }

//     return count
// }

//sort + 2 pointer
// func maxOperations(nums []int, k int) int {
//     sort.Ints(nums)

//     left, right := 0, len(nums)-1
//     count := 0

//     for left < right {
//         sum := nums[left] + nums[right]
//         if sum == k {
//             count++
//             left++
//             right--
//         } else if sum < k {
//             left++
//         } else {
//             right--
//         }
//     }

//     return count
// }

// Hashmap
func maxOperations(nums []int, k int) int {
	freq := make(map[int]int)
	count := 0

	for _, ele := range nums {
		cmp := k - ele
		if freq[cmp] > 0 {
			count++
			freq[cmp]--
		} else {
			freq[ele]++
		}
	}

	return count
}
