package main

import "fmt"

type stackNode struct {
	payload int
	next    *stackNode
}

type stack struct {
	head *stackNode
}

type StackWithMin struct {
	valueStack, minStack stack
}

func (this *stack) Push(value int) {
	node := &stackNode{payload: value, next: this.head}
	this.head = node
}

func (this *stack) Pop() int {
	if this.head == nil {
		return -1
	}

	result := this.head.payload
	this.head = this.head.next
	return result
}

func (this *stack) Peek() int {
	if this.head == nil {
		return -1
	}
	return this.head.payload
}

func (this *StackWithMin) Push(value int) {
	if this.minStack.head == nil || value < this.minStack.Peek() {
		this.valueStack.Push(value)
		this.minStack.Push(value)
	} else {
		this.valueStack.Push(value)
	}
}

func (this *StackWithMin) Pop() int {
	result := this.valueStack.Pop()
	if result == this.minStack.Peek() {
		this.minStack.Pop()
	}
	return result
}

func (this *StackWithMin) Min() int {
	return this.minStack.Peek()
}

func main() {
	s := new(StackWithMin)
	s.Push(1)
	fmt.Println(s.Min())
	s.Push(5)
	s.Push(10)
	fmt.Println(s.Min())
	s.Pop()
	s.Pop()
	s.Pop()
	s.Push(10)
	fmt.Println(s.Min())
	s.Push(15)
	fmt.Println(s.Min())
	s.Push(5)
	fmt.Println(s.Min())
	s.Pop()
	fmt.Println(s.Min())

}
