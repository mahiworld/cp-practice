package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func pairSum(head *ListNode) int {

//     //find middle element
//     slow := head
//     fast := head
//     for fast!= nil && fast.Next != nil {
//         slow = slow.Next
//         fast = fast.Next.Next
//     }

//     //reverse 2nd half of list
//     var prev *ListNode
//     for slow != nil{
//         temp := slow.Next
//         slow.Next = prev
//         prev = slow
//         slow = temp
//     }

//     //traverse two ll and find max twin
//     maxSum := 0
//     for prev != nil{
//         twinSum := head.Val + prev.Val
//         if twinSum > maxSum{
//             maxSum = twinSum
//         }

//         head = head.Next
//         prev = prev.Next
//     }

//     return maxSum
// }

func pairSum(head *ListNode) int {
	// Find the middle element and reverse the first half of the list
	var prev *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next

		// Reverse the first half of the list while moving slow
		temp := slow.Next
		slow.Next = prev
		prev = slow
		slow = temp
	}

	// Handle odd-length lists by skipping the middle element
	if fast != nil {
		slow = slow.Next
	}

	// Calculate the maximum twin sum
	maxSum := 0
	for prev != nil {
		twinSum := prev.Val + slow.Val
		if twinSum > maxSum {
			maxSum = twinSum
		}
		prev = prev.Next
		slow = slow.Next
	}

	return maxSum
}
