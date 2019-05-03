package main

import (
	"fmt"
)

func getNext(n uint) uint {
	c := n
	var c1 uint = 0
	var c0 uint = 0

	for ; c > 0 && c&1 == 0; c = c >> 1 {
		c0++
	}
	for ; c&1 == 1; c = c >> 1 {
		c1++
	}
	fmt.Printf("c0: %v, c1: %v\n", c0, c1)
	p := c0 + c1
	fmt.Printf("%b\n", n)
	n = n | 1<<p
	fmt.Printf("%b\n", n)
	n = n & ^((1 << p) - 1)
	fmt.Printf("%b\n", n)
	n = n | (1<<(c1-1) - 1)
	fmt.Printf("%b\n", n)
	return n
}

func getPrev(n uint) uint {
	c := n
	var c0 uint = 0
	var c1 uint = 0

	for ; c&1 == 1 && c > 0; c = c >> 1 {
		c1++
	}

	for ; c&1 == 0 && c > 0; c = c >> 1 {
		c0++
	}
	fmt.Printf("c0: %v, c1: %v\n", c0, c1)
	p := c1 + c0
	fmt.Printf("%b\n", n)
	n = n ^ (1 << p)
	fmt.Printf("%b\n", n)
	n = n & ^((1 << p) - 1)
	fmt.Printf("%b\n", n)
	n = n | ((1<<(c1+1) - 1) << (c0 - 1))
	fmt.Printf("%b\n", n)
	return n
}

func main() {
	//fmt.Println(getNext(13948))
	fmt.Println(getPrev(10115))

}
