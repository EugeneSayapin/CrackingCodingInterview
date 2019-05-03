package main

import (
	"fmt"
)

type node struct {
	name       string
	dependsOn  map[string]bool
	dependants map[string]bool
}

func buildNodes(projects []string, dependecies [][]string) map[string]*node {
	var result = make(map[string]*node)
	for _, project := range projects {
		node := &node{name: project, dependsOn: make(map[string]bool), dependants: make(map[string]bool)}
		result[project] = node
	}
	for _, dependency := range dependecies {
		result[dependency[1]].dependsOn[dependency[0]] = true
		result[dependency[0]].dependants[dependency[1]] = true
	}
	return result
}

func getBuildOrder(nodes map[string]*node) []string {
	var result = make([]string, len(nodes))
	var endOfList = addDepenciesNodes(nodes, result, 0)
	for i := 0; i < len(result); i++ {

		// for _, p := range result {
		// 	fmt.Printf("%v, ", p)
		// }
		// fmt.Println()

		node := nodes[result[i]]
		//fmt.Printf("processing node: %v\n", node.name)
		for key := range node.dependants {
			dependant := nodes[key]
			//fmt.Printf("processing dependant: %v\n", dependant.name)
			delete(dependant.dependsOn, result[i])
			if len(dependant.dependsOn) == 0 {
				result[endOfList] = dependant.name
				endOfList++
			}
		}
	}
	return result
}

func addDepenciesNodes(nodes map[string]*node, array []string, startIndex int) int {
	var found = 0
	for _, node := range nodes {
		if len(node.dependsOn) == 0 {
			array[startIndex+found] = node.name
			found++
		}
	}
	return found

}

// func printNodes(nodes map[string]*node) {
// 	for _, node := range nodes {
// 		fmt.Printf("Project Name: %v\nDependencies: ", node.name)
// 		for dep := range node.dependsOn {
// 			fmt.Printf("%v, ", dep)
// 		}
// 		fmt.Println()
// 	}
// }

func main() {
	var nodes = buildNodes([]string{"a", "b", "c", "d", "e", "f", "g"}, [][]string{{"f", "c"}, {"f", "b"}, {"f", "a"},
		{"c", "a"}, {"b", "a"}, {"b", "e"}, {"a", "e"}, {"d", "g"}})

	//printNodes(nodes)
	var result = getBuildOrder(nodes)
	for _, p := range result {
		fmt.Printf("%v, ", p)
	}
	fmt.Println()
	var nodes1 = buildNodes([]string{"a", "b", "c", "d", "e", "f"}, [][]string{{"a", "d"}, {"f", "b"}, {"d", "b"},
		{"f", "a"}, {"d", "c"}})

	//printNodes(nodes)
	var result1 = getBuildOrder(nodes1)
	for _, p := range result1 {
		fmt.Printf("%v, ", p)
	}

}
