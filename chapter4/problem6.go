package main

import (
	"fmt"
)

type treeNode struct {
	payload int
	left    *treeNode
	right   *treeNode
	parent  *treeNode
}

func getSuccessor(node *treeNode) int {
	if node.right != nil {
		return leftMost(node.right)
	}
	current := node
	parent := node.parent
	for parent.left != current {
		current = parent
		parent = current.parent
	}
	return parent.payload
}

func leftMost(node *treeNode) int {

	current := node
	for ; current.left != nil; current = current.left {
		//fmt.Println(current.payload)
	}
	return current.payload
}

func main() {
	n3 := &treeNode{payload: 3}
	n7 := &treeNode{payload: 7}
	n5 := &treeNode{payload: 5, left: n3, right: n7}
	n3.parent = n5
	n7.parent = n5
	n17 := &treeNode{payload: 17}
	n15 := &treeNode{payload: 15, right: n17}
	n17.parent = n15
	n10 := &treeNode{payload: 10, left: n5, right: n15}
	n5.parent = n10
	n15.parent = n10
	n30 := &treeNode{payload: 30}
	n20 := &treeNode{payload: 20, left: n10, right: n30}
	n10.parent = n20
	n30.parent = n20
	fmt.Println(getSuccessor(n7))
}
