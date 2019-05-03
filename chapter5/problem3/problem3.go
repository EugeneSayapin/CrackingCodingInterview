package main

import (
	"fmt"
)

func findFlip(input uint) int {
	current := 0
	prev := 0
	biggestPos := -1
	biggestSeq := 0
	for i := 0; input > 0; input = input >> 1 {
		if input&1 == 1 {
			current++
		} else {
			if input&2 > 0 {
				if (prev + current) > biggestSeq {
					biggestSeq = prev + current
					biggestPos = i
				}
				prev = current
				current = 0
			} else {
				prev = 0
				current = 0
			}
		}
		i++
	}
	return biggestPos
}

func main() {
	fmt.Println(findFlip(1775))

}
