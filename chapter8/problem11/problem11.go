package main

import (
	"fmt"
)

type intPair struct {
	amount int
	index  int
}

func countWays(amount int, demons []int) int {
	return countWaysIntenal(amount, demons, 0, map[intPair]int{})
}

func countWaysIntenal(amount int, demons []int, index int, calculatedAmounts map[intPair]int) int {
	fmt.Printf("count ways int. amount: %v, index: %v\n", amount, index)
	pair := intPair{amount: amount, index: index}
	value, exists := calculatedAmounts[pair]
	if exists {
		return value
	}
	if demons[index] == 1 {
		return 1
	}
	ways := 0

	for i := 0; amount >= i*demons[index]; i++ {
		ways += countWaysIntenal(amount-i*demons[index], demons, index+1, calculatedAmounts)
	}
	calculatedAmounts[pair] = ways
	return ways

}

func main() {
	fmt.Println(countWays(100, []int{25, 10, 5, 1}))
}
