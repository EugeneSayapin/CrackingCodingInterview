package main
import (
	"fmt"
	"math"
)

type treeNode struct {
	payload int
	left *treeNode
	right *treeNode
}
var lastSeenValue = math.MinInt32

func checkBST(node *treeNode) bool {
	if node == nil {
		return true
	}
	if !checkBST(node.left) {
		return false
	}

	if lastSeenValue > node.payload {
		return false
	}
	lastSeenValue = node.payload

	if !checkBST(node.right) {
		return false
	}
	return true
}


func main(){
	//fmt.Println(lastSeenValue)

	n3 := &treeNode {payload : 3}
	n7 := &treeNode {payload : 7}
	n5 := &treeNode {payload : 5, left : n3, right : n7}
	n17 := &treeNode {payload : 17}
	n15 := &treeNode {payload : 15, right : n17}
	n10 := &treeNode {payload : 10, left : n5, right : n15}
	n30 := &treeNode {payload :30}
	n20 := &treeNode {payload : 20, left: n10, right:n30}
	fmt.Println(checkBST(n20))
}