package main

// two pointers

//Approach:1 - shift & fill
// func moveZeroes(nums []int) {

//     index := 0

//     for _, num := range nums{
//         if num != 0{
//             nums[index] = num
//             index++
//         }
//     }

//     for i:=index; i<len(nums); i++{
//         nums[i] = 0
//     }
// }

// Appraoch:2 - Two pointer swap
func moveZeroes(nums []int) {

	l := 0
	for r := 0; r < len(nums); r++ {
		if nums[r] == 0 {
			continue
		}

		nums[l], nums[r] = nums[r], nums[l]
		l++
	}
}
