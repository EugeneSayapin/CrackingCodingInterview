package main

import (
	"fmt"
)

type treeNode struct {
	payload int
	left    *treeNode
	right   *treeNode
}

var runningSums map[int]int
var result = 0

func countPaths(node *treeNode, targetSum int, runningSum int) {
	if node == nil {
		return
	}

	runningSum += node.payload
	runningSums[runningSum]++
	result += runningSums[runningSum-targetSum]
	countPaths(node.left, targetSum, runningSum)
	countPaths(node.right, targetSum, runningSum)
	runningSums[runningSum]--

}

func main() {
	runningSums = make(map[int]int)
	n3leaf := &treeNode{payload: 3}
	n_2 := &treeNode{payload: -2}
	n3 := &treeNode{payload: 3, left: n3leaf, right: n_2}
	n1 := &treeNode{payload: 1}
	n2 := &treeNode{payload: 2, right: n1}
	n5 := &treeNode{payload: 5, right: n2, left: n3}
	n11 := &treeNode{payload: 11}
	n_3 := &treeNode{payload: -3, right: n11}
	n10 := &treeNode{payload: 10, left: n5, right: n_3}
	countPaths(n10, 8, 0)
	fmt.Println(result)
}
