package main

import (
	"fmt"
)

func findMagicIndex(array []int, left int, right int) int {
	fmt.Printf("left: %v, right: %v\n", left, right)

	if left < 0 || right >= len(array) || right-left < 0 {
		return -1
	}
	mid := (left + right) / 2
	if array[mid] == mid {
		return mid
	}
	leftBoundary := mid - 1
	if array[mid] < mid {
		leftBoundary = array[mid]
	}

	leftResult := findMagicIndex(array, 0, leftBoundary)
	if leftResult > 0 {
		return leftResult
	}
	return findMagicIndex(array, mid+1, right)

}

func main() {
	fmt.Println(findMagicIndex([]int{-10, -5, 2, 2, 2, 3, 4, 7, 9, 12, 13}, 0, 10))
}
