package main

//Find Peak Element
// func findPeakElement(nums []int) int {
//     l, r := 0, len(nums)-1

//     for l < r {
//         mid := l + (r-l)/2

//         if nums[mid] > nums[mid+1] {
//             r = mid
//         } else {
//             l = mid + 1
//         }
//     }

//     return l
// }

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2

		if (mid == 0 || nums[mid] > nums[mid-1]) && (mid == len(nums)-1 || nums[mid] > nums[mid+1]) {
			return mid
		}

		if mid > 0 && nums[mid] > nums[mid+1] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}
