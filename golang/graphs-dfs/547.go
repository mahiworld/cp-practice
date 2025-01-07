package main

// Number of Provinces
// Start from any node and apply dfs for that node and increment count
// continue for all unvisited vertices
func findCircleNum(isConnected [][]int) int {
	vertices := len(isConnected)
	visited := make([]bool, vertices)
	count := 0

	for i := 0; i < vertices; i++ {
		if !visited[i] {
			dfs(isConnected, i, visited)
			count++
		}
	}

	return count
}

func dfs(matrix [][]int, v int, visited []bool) {
	visited[v] = true
	n := len(matrix)

	for i := 0; i < n; i++ {
		if matrix[v][i] == 1 && !visited[i] {
			dfs(matrix, i, visited)
		}
	}
}
