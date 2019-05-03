package main

import (
	"fmt"
)

var ways []int

func countWays(n int) int {
	fmt.Printf("n: %v\n", n)
	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1
	}
	fmt.Printf("ways[n]: %v\n", ways[n])
	if ways[n] != 0 {
		return ways[n]
	}

	result := countWays(n-3) + countWays(n-2) + countWays(n-1)
	ways[n] = result
	return result

}

func main() {
	n := 15
	ways = make([]int, n+1)
	fmt.Println(countWays(n))
}
