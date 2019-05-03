package main

import "fmt"

type stackNode struct {
	payload int
	next    *stackNode
}

type stack struct {
	top *stackNode
}

func (this *stack) push(value int) {
	node := &stackNode{payload: value, next: this.top}
	this.top = node
}

func (this *stack) pop() int {
	result := this.top.payload
	this.top = this.top.next
	return result
}

func (this *stack) isEmpty() bool {
	return this.top == nil
}

func (this stack) peek() int {
	return this.top.payload
}

func sort(orig *stack) {
	temp := &stack{}
	var current int
	for !orig.isEmpty() {
		current = orig.pop()
		if temp.isEmpty() || current > temp.peek() {
			temp.push(current)
		} else {

			for !temp.isEmpty() && temp.peek() > current {
				orig.push(temp.pop())
			}
			orig.push(current)
		}
	}
	for !temp.isEmpty() {
		orig.push(temp.pop())
	}

}
func main() {
	orig := &stack{}

	orig.push(8)
	orig.push(12)
	orig.push(1)
	orig.push(3)
	orig.push(5)
	orig.push(10)
	orig.push(7)

	sort(orig)

	for !orig.isEmpty() {
		fmt.Println(orig.pop())
	}
}
