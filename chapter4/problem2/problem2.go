package main

import (
	"fmt"
)

type node struct {
	value int
	left  *node
	right *node
}

func buildBST(array []int, start int, end int) *node {
	if start > end {
		return nil
	}
	if start == end {
		return &node{value: array[start]}
	}

	mid := (start + end) / 2

	result := &node{value: array[mid], left: buildBST(array, start, mid-1), right: buildBST(array, mid+1, end)}
	return result
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	result := buildBST(array, 0, 5)
	fmt.Print(result)
}
