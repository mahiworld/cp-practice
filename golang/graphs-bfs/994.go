package main

//Rotting Oranges

func orangesRotting(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	queue := [][]int{}
	freshCount := 0
	minutes := 0

	// Collect initial rotten oranges and count fresh oranges
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j, 0}) // Add rotten orange to queue with time 0
			} else if grid[i][j] == 1 {
				freshCount++
			}
		}
	}

	// If there are no fresh oranges, return 0
	if freshCount == 0 {
		return 0
	}

	// BFS to spread rotting effect
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		row, col, time := node[0], node[1], node[2]
		minutes = time

		for _, dir := range directions {
			newRow, newCol := row+dir[0], col+dir[1]
			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && grid[newRow][newCol] == 1 {
				grid[newRow][newCol] = 2 // Rotten the fresh orange
				freshCount--             // Reduce fresh count
				queue = append(queue, []int{newRow, newCol, time + 1})
			}
		}
	}

	// If fresh oranges remain, return -1
	if freshCount > 0 {
		return -1
	}

	return minutes
}
