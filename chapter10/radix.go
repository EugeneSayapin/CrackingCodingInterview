package main

import (
	"fmt"
	"math"
)

const base int = 10

func sort(array *[]int) {
	max := math.MaxInt32
	for i := 0; max > 0; i++ {
		max /= base
		*array = countSort(*array, i)
	}

}

func countSort(array []int, position int) []int {
	counts := make([]int, base)
	divider := 1
	for i := 0; i < position; i++ {
		divider *= base
	}
	for _, i := range array {
		counts[(i/divider)%base]++
	}
	for i := 1; i <= base-1; i++ {
		counts[i] += counts[i-1]
	}
	for i := base - 1; i > 0; i-- {
		counts[i] = counts[i-1]
	}
	counts[0] = 0
	result := make([]int, len(array))

	for _, i := range array {
		index := (i / divider) % base
		result[counts[index]] = i
		counts[index]++
	}

	return result

}

func main() {
	array := []int{633, 115, 986, 13, 9, 100, 67, 55}
	sort(&array)
	for _, i := range array {
		fmt.Printf("%v, ", i)
	}
}
