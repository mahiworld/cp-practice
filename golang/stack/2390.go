package main

//Approach : 1
// func removeStars(s string) string {
//     stack := make([]rune, 0)

//     for _, ele := range s{
//         if string(ele) == "*"{
//             pop(&stack)
//             continue
//         }

//         push(&stack, ele)
//     }

//     return string(stack)
// }

// func push(stack *[]rune, ele rune) {
//     *stack = append(*stack, ele)
// }

// func pop(stack *[]rune){
//     lenth := len(*stack)
//     if lenth <=0{
//         return
//     }
//     *stack = (*stack)[:lenth-1]
// }

// Approach:1 optimization
func removeStars(s string) string {
	stack := make([]rune, 0)

	for _, ele := range s {
		if ele == '*' {
			// Remove the last element
			stack = stack[:len(stack)-1]
		} else {
			// Append the current character
			stack = append(stack, ele)
		}
	}

	return string(stack)
}

// // Approach:2 In-place
// func removeStars(s string) string{
//     stack := []rune(s) // in golang because we can access string directly (bcoz strings in Go are immutable)
//     j := 0
//     for _, ele := range stack {
//         if ele == '*' {
//             if j > 0 {
//                 j--
//             }
//         } else {
//             stack[j] = ele
//             j++
//         }
//     }

//     return string(stack[:j])
// }
