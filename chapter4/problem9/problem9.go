package main

import (
	"fmt"
)

type node struct {
	payload int
	left    *node
	right   *node
}

func allSequences(root *node) [][]int {

	result := make([][]int, 0)
	if root == nil {
		result = append(result, make([]int, 0))
		return result
	}

	fmt.Printf("allSec\nnode: %v\n", root.payload)

	prefix := []int{root.payload}
	leftSeq := allSequences(root.left)
	rightSeq := allSequences(root.right)

	for _, left := range leftSeq {
		for _, right := range rightSeq {
			weaved := make([][]int, 0)
			weaveLists(left, right, &weaved, prefix)
			for _, weavedElem := range weaved {
				result = append(result, weavedElem)
			}
		}
	}

	return result
}

func weaveLists(first []int, second []int, results *[][]int, prefix []int) {

	// fmt.Println("weave")
	// fmt.Println("first")
	// printList(first)
	// fmt.Println("second")
	// printList(second)
	// fmt.Println("prefix")
	// printList(prefix)
	// fmt.Printf("\nresults len:%v\n", len(*results))
	// fmt.Println("weave end")

	if len(first) == 0 || len(second) == 0 {
		result := make([]int, len(prefix))
		copy(result, prefix)

		for _, firstElem := range first {
			result = append(result, firstElem)
		}

		for _, secondElem := range second {
			result = append(result, secondElem)
		}
		// fmt.Printf("appening to results")
		*results = append(*results, result)
		// fmt.Printf("\n new results len:%v\n", len(*results))
		return
	}

	headFirst := first[0]
	prefix = append(prefix, headFirst)
	weaveLists(first[1:], second, results, prefix)
	prefix = prefix[:len(prefix)-1]

	headSecond := second[0]
	prefix = append(prefix, headSecond)
	weaveLists(first, second[1:], results, prefix)
	prefix = prefix[:len(prefix)-1]

}

func printList(list []int) {
	fmt.Printf("{")
	for _, item := range list {
		fmt.Printf("%v, ", item)
	}
	fmt.Printf("}")
}

func main() {
	// n1 := &node{payload: 1}
	// n3 := &node{payload: 3}
	// n2 := &node{payload: 2, left: n1, right: n3}

	n5 := &node{payload: 5}
	n15 := &node{payload: 15}
	n10 := &node{payload: 10, left: n5, right: n15}
	n25 := &node{payload: 25}
	n20 := &node{payload: 20, left: n10, right: n25}
	// n65 := &node{payload: 65}
	// n80 := &node{payload: 80}
	// n70 := &node{payload: 70, left: n65, right: n80}
	// n60 := &node{payload: 60, right: n70}
	// n50 := &node{payload: 50, left: n20, right: n60}

	result := allSequences(n20)

	for _, array := range result {
		printList(array)
	}

	// array1 := []int{1, 2}
	// array2 := []int{3, 4}
	// results := [][]int{}
	// weaveLists(array1, array2, &results, []int{})
	// for _, array := range results {
	// 	printList(array)
	// }
}
