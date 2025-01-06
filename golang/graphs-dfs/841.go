package main

//Keys and Rooms

// func canVisitAllRooms(rooms [][]int) bool {
//     visited := map[int]bool{0:true}
//     keys := rooms[0]

//     for len(keys)>0{
//         key := keys[0]
//         keys = keys[1:]

//         if visited[key]{
//             continue
//         }else{
//             visited[key] = true
//             keys = append(keys, rooms[key]...)
//         }

//         if len(rooms) == len(visited){
//             return true
//         }
//     }

//     return false
// }

func canVisitAllRooms(rooms [][]int) bool {
	visited := map[int]bool{0: true}
	stack := rooms[0]

	for len(stack) > 0 {
		key := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[key] {
			visited[key] = true
			stack = append(stack, rooms[key]...)
		}
	}

	return len(rooms) == len(visited)
}
