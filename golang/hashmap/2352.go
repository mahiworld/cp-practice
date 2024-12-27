package main

import (
	"strconv"
	"strings"
)

// func equalPairs(grid [][]int) int {
//     count := 0
//     rows := len(grid)
//     cols := 0
//     if rows > 0{
//         cols = len(grid[0])
//     }

//     for i:=0; i< rows; i++{
//         for j:=0;j<rows;j++{
// match := true
//             for k:=0;k<cols;k++{
//                 if grid[i][k] != grid[k][j]{
// match = false
//                     break
//                 }
//                 if match{
//                     count++
//                 }
//             }
//         }
//     }

//     return count
// }

func equalPairs(grid [][]int) int {

	count := 0

	n := len(grid)

	rowMap := make(map[string]int)
	// colMap := make(map[string]int)

	for i := 0; i < n; i++ {
		rowMap[toString(grid[i])]++
	}

	for i := 0; i < n; i++ {
		strCol := make([]string, 0)
		for j := 0; j < n; j++ {
			strCol = append(strCol, strconv.Itoa(grid[j][i]))
		}

		count += rowMap[strings.Join(strCol, " ")]

		// colMap[strings.Join(strCol, " ")]++
	}

	// for key, val := range rowMap{
	//     if colMap[key]>0{
	//         count += val * colMap[key]
	//     }
	// }

	return count
}

func toString(nums []int) string {
	strNums := make([]string, len(nums))
	for i, num := range nums {
		strNums[i] = strconv.Itoa(num)
	}

	return strings.Join(strNums, " ")
}
