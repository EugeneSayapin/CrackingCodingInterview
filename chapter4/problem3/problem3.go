package main

type node struct {
	value int
	left  *node
	right *node
}

type listNode struct {
	payload int
	next    *listNode
}

type listOfListNode struct {
	payload *listNode
	next    *listOfListNode
}

func addToList(value int, level int, head *listOfListNode) {
	current := head
	var prev *listOfListNode
	for i := 0; i <= level; i++ {
		if current == nil {
			newOne := &listOfListNode{}
			prev.next = newOne
			current = newOne
		}
		if i == level {
			val := &listNode{value, current.payload}
			current.payload = val
		}
		prev = current
		current = current.next
	}
}

func getLevels(root *node, levels *listOfListNode, level int) {
	if root == nil {
		return
	}
	addToList(root.value, level, levels)
	getLevels(root.left, levels, level+1)
	getLevels(root.right, levels, level+1)
}

func main() {
	node1 := &node{1, nil, nil}
	node2 := &node{2, nil, nil}
	node3 := &node{3, node1, node2}
	levels := &listOfListNode{}
	getLevels(node3, levels, 0)

}
