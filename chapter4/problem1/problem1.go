package main

import "fmt"

type node struct {
	visited  bool
	adjacent []*node
}

type listNode struct {
	payload *node
	next    *listNode
}
type queue struct {
	head *listNode
	tail *listNode
}

func (this *queue) enqueue(node *node) {
	listNode := &listNode{payload: node, next: nil}
	if this.tail == nil && this.head == nil {
		this.head = listNode
		this.tail = listNode
		return
	}
	this.tail.next = listNode
	this.tail = listNode
}

func (this *queue) dequeue() *node {
	result := this.head.payload
	this.head = this.head.next
	if this.head == nil {
		this.tail = nil
	}
	return result
}

func search(start *node, target *node) bool {
	queue := &queue{}
	queue.enqueue(start)
	for queue.head != nil {
		current := queue.dequeue()
		if current == target {
			return true
		}
		current.visited = true
		for _, child := range current.adjacent {
			queue.enqueue(child)
		}
	}
	return false
}

func main() {

	n1 := &node{}
	n2 := &node{}
	n3 := &node{}
	n4 := &node{}
	n5 := &node{}
	n6 := &node{}
	n7 := &node{}
	n1.adjacent = []*node{n3, n4}
	n2.adjacent = []*node{n1}
	n3.adjacent = []*node{n2}
	n4.adjacent = []*node{n5, n6}
	n5.adjacent = []*node{}
	n6.adjacent = []*node{n7}
	n7.adjacent = []*node{}

	fmt.Println(search(n5, n7))
}
