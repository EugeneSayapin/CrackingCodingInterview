package main

import "fmt"

func rotate(array [][]int) {
	n := len(array)
	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - layer - 1
		for i := first; i < last; i++ {
			temp := array[first+i][last]
			array[first+i][last] = array[first][first+i]
			array[first][first+i] = array[last-i][first]
			array[last-i][first] = array[last][last-i]
			array[last][last-i] = temp
		}
	}
}

func main() {
	array1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(array1)
	fmt.Println(array1)

	array2 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	rotate(array2)
	fmt.Println(array2)
}
