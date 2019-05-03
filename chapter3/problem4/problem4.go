package main

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

type queue struct {
	simple   *stack
	resersed *stack
}

func (this *queue) Enqueue(value int) {
	this.simple.push(value)
}

func (this *queue) Dequeue() int {
	if !this.resersed.isEmpty() {
		return this.resersed.pop()
	}
	for !this.simple.isEmpty() {
		this.resersed.push(this.simple.pop())
	}
	return this.resersed.pop()
}

func main() {
	queue := new(queue)
	simpleStack := stack{}
	reversedStack := stack{}
	queue.resersed = &reversedStack
	queue.simple = &simpleStack

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Dequeue()
	queue.Enqueue(3)
	queue.Dequeue()
	queue.Dequeue()

}
