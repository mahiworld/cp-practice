package main

// Nearest Exit from Entrance in Maze
func nearestExit(maze [][]byte, entrance []int) int {
	rows, cols := len(maze), len(maze[0])
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	queue := [][]int{{entrance[0], entrance[1], 0}}
	maze[entrance[0]][entrance[1]] = '+'

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		row, col, stepCount := node[0], node[1], node[2]

		for _, dir := range directions {
			newRow, newCol := row+dir[0], col+dir[1]
			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && maze[newRow][newCol] == '.' {
				if newRow == 0 || newRow == rows-1 || newCol == 0 || newCol == cols-1 {
					return stepCount + 1
				}

				maze[newRow][newCol] = '+'
				queue = append(queue, []int{newRow, newCol, stepCount + 1})
			}
		}
	}

	return -1
}
