package main

import (
	"fmt"
)

func countDifferentBits(n1, n2 uint) int {
	c := n1 ^ n2
	result := 0

	for ; c > 0; c = c & (c - 1) {
		result++
	}
	return result

}

func main() {
	fmt.Println(countDifferentBits(29, 15))

}
