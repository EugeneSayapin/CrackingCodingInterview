package main;

import "fmt"

type treeNode struct {
	payload int
	left *treeNode
	right *treeNode
}

func getHeight(root *treeNode) int {
	if root == nil {
		return 0
	}
	var leftHeight = getHeight(root.left)
	if leftHeight == -1 {
		return -1
	}
	var rightHeight = getHeight(root.right)
	if rightHeight == -1 {
		return -1
	}
	fmt.Printf("payload: %v; leftHeight: %v; rightHeight: %v\n", root.payload, leftHeight, rightHeight)
	var heightDif = 0;
	if rightHeight > leftHeight {
		heightDif = rightHeight - leftHeight
	} else {
		heightDif = leftHeight - rightHeight
	}

	if heightDif > 1 {
		return -1
	} else {
		if leftHeight > rightHeight {
			return leftHeight +1
		} else {
			return rightHeight +1
		}

	}
}

func main() {
	//0 := &treeNode{payload : 0}
	l1 := &treeNode{payload : 1}
	l2 := &treeNode{payload : 2, left : l1}
	r2 := &treeNode{payload : 3}
	l3 := &treeNode{payload : 4, left: l2, right: r2}
	fmt.Println(getHeight(l3))
}



