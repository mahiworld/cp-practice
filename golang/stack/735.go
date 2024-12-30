package main

//Stack approach
// type Ast struct{
//     Val int
//     Dir bool // true means + (right), false means - (left)
// }

// func asteroidCollision(asteroids []int) []int {
//     stack := make([]Ast, 0)

//     for _, ele := range asteroids{
//         newEle := Ast{
//             Val: absInt(ele),
//             Dir: ele > 0,
//         }

//         for len(stack)>0 && stack[len(stack)-1].Dir && !newEle.Dir{
//             top := stack[len(stack)-1]

//             if top.Val == newEle.Val {//Both Exploded
//                 stack = stack[:len(stack)-1]
//                 newEle = Ast{Val: 0, Dir: true}
//                 break
//             } else if top.Val > newEle.Val {//new exploded, top servive
//                 newEle = Ast{Val: 0, Dir: true}
//                 break
//             } else {//top exploded, new servive
//                 stack = stack[:len(stack)-1]
//             }
//         }

//         // Push the new asteroid if it hasn't exploded
//         if newEle.Val != 0 {
//             stack = append(stack, newEle)
//         }
//     }

//     // Convert stack to result
//     resp := make([]int, 0, len(stack))
//     for _, ele := range stack {
//         if ele.Dir {
//             resp = append(resp, ele.Val)
//         } else {
//             resp = append(resp, -ele.Val)
//         }
//     }

//     return resp
// }

func absInt(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

// Optimized stack approach
func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)
	for _, ele := range asteroids {

		for len(stack) > 0 && stack[len(stack)-1] > 0 && ele < 0 {
			top := stack[len(stack)-1]

			if absInt(top) == absInt(ele) { //both explode
				stack = stack[:len(stack)-1]
				ele = 0
				break
			} else if absInt(top) > absInt(ele) { //new explode
				ele = 0
				break
			} else { //top explode
				stack = stack[:len(stack)-1]
			}
		}

		if ele != 0 {
			stack = append(stack, ele)
		}
	}

	return stack
}
