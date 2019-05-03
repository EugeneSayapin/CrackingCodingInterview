package main

import (
	"fmt"
)

type point struct {
	row    int
	column int
}

var visitedPoints map[point]bool

func getPath(position point, maze [][]bool, currentPath []point) []point {

	if visitedPoints[position] {
		return nil
	}
	visitedPoints[position] = true

	if (position.row == len(maze)-1) && (position.column == len(maze[0])) {
		return currentPath
	}
	if position.row >= len(maze) {
		return nil
	}
	if position.column >= len(maze[0]) {
		return nil
	}
	if maze[position.row][position.column] {
		return nil
	}
	stepDown := getPath(point{row: position.row + 1, column: position.column}, maze, append(currentPath, position))
	if stepDown != nil {
		return stepDown
	}
	return getPath(point{row: position.row, column: position.column + 1}, maze, append(currentPath, position))

}

func main() {
	visitedPoints = make(map[point]bool)
	maze := [][]bool{
		{false, false, false},
		{true, false, false},
		{false, false, false},
	}

	result := getPath(point{0, 0}, maze, []point{})

	for _, c := range result {
		fmt.Printf("{%v, %v}, ", c.row, c.column)
	}

}
