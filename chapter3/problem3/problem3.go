package main

type stackNode struct {
	payload int
	next    *stackNode
}

type stack struct {
	top   *stackNode
	count int
}

func (this *stack) push(value int) {
	node := &stackNode{payload: value, next: this.top}
	this.top = node
	this.count++
}

func (this *stack) pop() int {
	result := this.top.payload
	this.top = this.top.next
	this.count--
	return result
}

type arrayOfStacks struct {
	stacks            [5]*stack
	maxStackSize      int
	currentStackIndex int
}

func newArrayOfStacks() *arrayOfStacks {
	array := new(arrayOfStacks)
	array.maxStackSize = 5
	array.currentStackIndex = 0
	for index := range array.stacks {
		array.stacks[index] = new(stack)
	}
	return array
}

func (this *arrayOfStacks) push(payload int) {
	if this.stacks[this.currentStackIndex].count == this.maxStackSize {
		this.currentStackIndex++
	}
	this.stacks[this.currentStackIndex].push(payload)
}

func (this *arrayOfStacks) pop() int {
	if this.stacks[this.currentStackIndex].count == 0 {
		this.currentStackIndex--
	}
	return this.stacks[this.currentStackIndex].pop()
}

func main() {
	arrayOfStacks := newArrayOfStacks()
	arrayOfStacks.push(1)
	arrayOfStacks.push(2)
	arrayOfStacks.push(3)
	arrayOfStacks.push(4)
	arrayOfStacks.push(5)
	arrayOfStacks.push(6)
	arrayOfStacks.pop()
	arrayOfStacks.pop()
	arrayOfStacks.push(7)
	arrayOfStacks.push(8)

}
