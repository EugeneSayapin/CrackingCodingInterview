package main

import "fmt"

type MultiStack struct {
	values   []int
	sizes    []int
	capacity int
}

func (stack *MultiStack) Init(capacity, numberOfStack int) {
	stack.values = make([]int, capacity*numberOfStack)
	stack.sizes = make([]int, numberOfStack)
	stack.capacity = capacity
}

func (stack *MultiStack) Push(value, stackNumber int) {
	if stackNumber < 0 || stackNumber >= len(stack.sizes) {
		return
	}
	if stack.sizes[stackNumber] > stack.capacity {
		return
	}
	stack.sizes[stackNumber]++
	stack.values[stack.capacity*stackNumber+stack.sizes[stackNumber]] = value
}

func (stack *MultiStack) Pop(stackNumber int) int {
	if stackNumber < 0 || stackNumber >= len(stack.sizes) {
		return -1
	}

	if stack.sizes[stackNumber] == 0 {
		return -1
	}

	stack.sizes[stackNumber]--
	return stack.values[stack.capacity*stackNumber+stack.sizes[stackNumber]+1]
}

func main() {
	stacks := new(MultiStack)
	stacks.Init(10, 3)
	stacks.Push(1, 0)
	stacks.Push(2, 1)
	stacks.Push(3, 2)
	stacks.Push(4, 0)
	stacks.Push(5, 1)
	stacks.Push(6, 2)

	fmt.Println(stacks.Pop(1))
	fmt.Println(stacks.Pop(1))
}
