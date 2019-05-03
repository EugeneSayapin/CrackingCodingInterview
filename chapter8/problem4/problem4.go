package main

import (
	"fmt"
)

func getAllSubsets(set []string) [][]string {
	result := [][]string{}

	for i := 0; i < 1<<uint(len(set)); i++ {
		currentSebuset := []string{}
		for j := 0; j < len(set); j++ {
			if i&(1<<uint(j)) != 0 {
				currentSebuset = append(currentSebuset, set[j])
			}
		}
		result = append(result, currentSebuset)
	}
	return result
}

func main() {
	toPrint := getAllSubsets([]string{"a", "b", "c", "d"})
	for _, set := range toPrint {
		fmt.Print("{")
		for _, s := range set {
			fmt.Printf("%v, ", s)
		}
		fmt.Print("}\n")
	}
}
