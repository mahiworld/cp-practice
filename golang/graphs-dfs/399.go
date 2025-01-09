package main

// Evaluate Division
type VE struct {
	v string
	e float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	adjList := make(map[string][]VE)
	for index, eq := range equations {
		a, b, val := eq[0], eq[1], values[index]
		adjList[a] = append(adjList[a], VE{v: b, e: float64(val)})
		adjList[b] = append(adjList[b], VE{v: a, e: float64(1 / val)})
	}

	result := make([]float64, 0)
	for _, query := range queries {
		start, end := query[0], query[1]
		if adjList[start] == nil || adjList[end] == nil {
			result = append(result, -1.0)
		} else {
			visited := make(map[string]bool)
			res := dfsCalEQ(adjList, query, 1.0, visited)
			// res := dfs(adjList, start, end, 1.0, visited)
			result = append(result, res)
		}
	}

	return result
}

func dfsCalEQ(adjList map[string][]VE, query []string, acc float64, visited map[string]bool) float64 {
	if query[0] == query[1] {
		return acc
	}
	visited[query[0]] = true

	for _, edge := range adjList[query[0]] {
		if !visited[edge.v] {
			res := dfsCalEQ(adjList, []string{edge.v, query[1]}, acc*edge.e, visited)
			if res != -1.0 {
				return res
			}
		}
	}

	return -1.0
}

// func dfs(adjList map[string][]VE, current, target string, acc float64, visited map[string]bool) float64 {
// 	if current == target {
// 		return acc
// 	}
// 	visited[current] = true

// 	for _, edge := range adjList[current] {
// 		if !visited[edge.v] {
// 			res := dfs(adjList, edge.v, target, acc*edge.e, visited)
// 			if res != -1.0 {
// 				return res
// 			}
// 		}
// 	}

// 	return -1.0
// }
