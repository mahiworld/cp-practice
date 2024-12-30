package main

import (
	"strconv"
	"strings"
)

// func decodeString(s string) string {
//     stack := make([]string, 0)
//     for _, ele := range s{
//         if string(ele) != "]" {
//             // Push characters, including digits and '['
//             stack = append(stack, string(ele))
//         }else{
//             // Extract the string inside the brackets
//             tempStr := ""
//             for len(stack) > 0 && stack[len(stack)-1] != "[" {
//                 tempStr = stack[len(stack)-1] + tempStr
//                 stack = stack[:len(stack)-1]
//             }

//             // Pop the '['
//             stack = stack[:len(stack)-1]

//             // Extract the repeat count (which can be multi-digit)
//             numStr := ""
//             for len(stack) > 0 && stack[len(stack)-1] >= "0" && stack[len(stack)-1] <= "9" {
//                 numStr = stack[len(stack)-1] + numStr
//                 stack = stack[:len(stack)-1]
//             }
//             repeat, _ := strconv.Atoi(numStr)

//             // Repeat the string and push it back to the stack
//             expandedStr := strings.Repeat(tempStr, repeat)
//             stack = append(stack, expandedStr)

//         }
//     }

//     return strings.Join(stack, "")
// }

func decodeString(s string) string {
	stack := make([]rune, 0)

	for _, ele := range s {
		if ele != ']' {
			// Push character directly to the stack
			stack = append(stack, ele)
		} else {
			// Extract the string inside the brackets
			tempStr := make([]rune, 0)
			for len(stack) > 0 && stack[len(stack)-1] != '[' {
				tempStr = append([]rune{stack[len(stack)-1]}, tempStr...) //to avoid reversing
				stack = stack[:len(stack)-1]
			}

			// Pop the '['
			stack = stack[:len(stack)-1]

			// Extract the repeat count (multi-digit numbers supported)
			numStr := make([]rune, 0)
			for len(stack) > 0 && stack[len(stack)-1] >= '0' && stack[len(stack)-1] <= '9' {
				numStr = append([]rune{stack[len(stack)-1]}, numStr...)
				stack = stack[:len(stack)-1]
			}
			repeat, _ := strconv.Atoi(string(numStr))

			// Repeat the string and push it back to the stack
			repeatedStr := strings.Repeat(string(tempStr), repeat)
			stack = append(stack, []rune(repeatedStr)...)
		}
	}

	return string(stack)
}
