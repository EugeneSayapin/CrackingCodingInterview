package main

import "fmt"

type node struct {
	payload int
	left    *node
	right   *node
}

func findAnc(root *node, p int, q int) (int, bool) {
	if root == nil {
		return 0, false
	}

	if root.payload == p {
		return p, true
	}
	if root.payload == q {
		return q, true
	}

	lvalue, lresult := findAnc(root.left, p, q)
	rvalue, rresult := findAnc(root.right, p, q)

	if !lresult && !rresult {
		return 0, false
	}
	if lresult && rresult {
		return root.payload, true
	}

	if lresult {
		return lvalue, true
	}

	return rvalue, true

}

func main() {
	n3 := &node{payload: 3}
	n7 := &node{payload: 7}
	n5 := &node{payload: 5, left: n3, right: n7}
	n17 := &node{payload: 17}
	n15 := &node{payload: 15, right: n17}
	n10 := &node{payload: 10, left: n5, right: n15}
	n30 := &node{payload: 30}
	n20 := &node{payload: 20, left: n10, right: n30}

	acc, _ := findAnc(n20, 5, 30)
	fmt.Println(acc)
}
