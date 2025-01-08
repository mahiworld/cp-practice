package main

//Problem: Reorder Routes to Make All Paths Lead to the City Zero

/*************Approach:1 ***********************/

// func minReorder(n int, connections [][]int) int {
// 	graph := make([][]map[int]int, n)
// 	for _, conn := range connections {
// 		a, b := conn[0], conn[1]
// 		graph[a] = append(graph[a], map[int]int{b: 1})
// 		graph[b] = append(graph[b], map[int]int{a: 0})
// 	}

// 	visited := make([]bool, n)
// 	count := 0

// 	dfs(graph, 0, visited, &count)

// 	return count
// }

// func dfs(adjList [][]map[int]int, v int, visited []bool, count *int) {
// 	visited[v] = true

// 	for _, neighborMap := range adjList[v] {
// 		for neighbor, dir := range neighborMap {
// 			if !visited[neighbor] {
// 				if dir == 1 {
// 					(*count)++
// 				}
// 				dfs(adjList, neighbor, visited, count)
// 			}
// 		}
// 	}
// }

/*************Approach:2 ***********************/

// NeighborDir defines a structure with a neighbor and a direction
type NeighborDir struct {
	neighbor int
	dir      bool
}

func minReorder(n int, connections [][]int) int {
	graph := make([][]NeighborDir, n)

	for _, conn := range connections {
		a, b := conn[0], conn[1]
		graph[a] = append(graph[a], NeighborDir{b, true})
		graph[b] = append(graph[b], NeighborDir{a, false})
	}

	visited := make([]bool, n)
	count := 0

	dfsgraph(graph, 0, visited, &count)

	return count
}

func dfsgraph(graph [][]NeighborDir, city int, visited []bool, count *int) {
	visited[city] = true

	for _, edge := range graph[city] {
		if !visited[edge.neighbor] {
			if edge.dir {
				(*count)++
			}
			dfsgraph(graph, edge.neighbor, visited, count)
		}
	}
}
