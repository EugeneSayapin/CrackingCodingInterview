package main

import (
	"fmt"
)

const gridSize = 8

func findPlacement() [][]int {
	result := make([][]int, 0)
	findPlacementInternal(0, make([]int, gridSize), &result)
	return result
}

func findPlacementInternal(row int, columns []int, results *[][]int) {
	if row == gridSize {
		solution := make([]int, gridSize)
		copy(solution, columns)
		*results = append(*results, solution)
		return
	}
	for col := 0; col < gridSize; col++ {
		if validPlacement(columns, row, col) {
			columns[row] = col
			findPlacementInternal(row+1, columns, results)
		}
	}
}

func validPlacement(columns []int, row int, col int) bool {
	for r := 0; r < row; r++ {
		if columns[r] == col {
			return false
		}
		if abs(col, columns[r]) == abs(row, r) {
			return false
		}
	}
	return true
}
func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func printSolutions(columns []int) {
	for row := 0; row < gridSize; row++ {
		for col := 0; col < gridSize; col++ {
			if columns[row] == col {
				fmt.Print(" Q ")
			} else {
				fmt.Print(" X ")
			}
		}
		fmt.Println("\n---------------------------")
	}
}

func main() {
	solutions := findPlacement()
	for _, solution := range solutions {
		fmt.Println("Next solution")
		printSolutions(solution)
	}
}
