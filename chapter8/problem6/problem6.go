package main

import (
	"fmt"
)

type tower struct {
	disks      []int
	currentTop int
	name       string
}

func (tower *tower) init(n int, name string) {
	tower.disks = make([]int, n)
	tower.name = name
	tower.currentTop = -1
}
func (tower *tower) push(value int) {
	tower.currentTop++
	//fmt.Printf("push to %v, current position %v, value %v\n", tower.name, tower.currentTop, value)
	tower.disks[tower.currentTop] = value
}

func (tower *tower) pop() int {

	temp := tower.disks[tower.currentTop]
	tower.currentTop--
	//fmt.Printf("pop from %v, current position %v, currentValue %v\n", tower.name, tower.currentTop, temp)
	return temp

}

func moveDisks(n int, from *tower, to *tower, buffer *tower) {
	//fmt.Printf("function start:\nn: %v, from: %v, to: %v, buffer: %v\n", n, from.name, to.name, buffer.name)
	if n == 0 {
		return
	}
	moveDisks(n-1, from, buffer, to)
	fmt.Printf("moving disk %v from %v to %v\n", from.disks[from.currentTop], from.name, to.name)
	to.push(from.pop())
	moveDisks(n-1, buffer, to, from)

}

func main() {
	n := 4
	origin := &tower{}
	buffer := &tower{}
	target := &tower{}

	origin.init(n, "first")
	buffer.init(n, "second")
	target.init(n, "third")

	for i := n; i > 0; i-- {
		origin.push(i)
	}
	moveDisks(n, origin, target, buffer)

}
