package main

import (
	"fmt"
	"math/rand"
)

type treeNode struct {
	payload int
	left    *treeNode
	right   *treeNode
	size    int
}

func (root *treeNode) insert(value int) {
	//fmt.Printf("insert value: %v; root payload: %v\n", value, root.payload)
	if root.payload > value {
		if root.left == nil {
			newNode := &treeNode{payload: value, size: 1}
			root.left = newNode
		} else {
			root.left.insert(value)
		}

	} else {
		if root.right == nil {
			newNode := &treeNode{payload: value, size: 1}
			root.right = newNode
		} else {
			root.right.insert(value)
		}
	}
	root.size++
}

func (root *treeNode) delete(value int, decrement bool) {
	fmt.Printf("delete value: %v; root payload: %v\n", value, root.payload)
	if root.payload > value {
		if root.left.payload != value {
			root.left.delete(value, decrement)
			if decrement {
				root.size--
			}
			return
		}
		//root left == value
		if root.left.left == nil || root.left.right == nil {
			fmt.Print("deleting\n")
			if root.left.left == nil {
				fmt.Printf("deleting left\nroot value: %v, root.left value: %v\n", root.payload, root.left.payload)
				fmt.Printf("new node: %v\n", root.left.right)
				root.left = root.left.right
				fmt.Printf("after value is set %v\n", root.left)
			} else {
				fmt.Print("deleting right\n")
				root.left = root.left.left
			}
			if decrement {
				fmt.Print("decrement\n")
				root.size--
			}
			return
		}
		fmt.Printf("before findMin\n root.left: %v, root.left.right: %v\n", root.left.payload, root.left.right.payload)
		minVal := root.left.right.findMin().payload
		root.left.payload = minVal
		root.left.size--
		root.size--
		root.left.delete(minVal, false)

		return
	} else {
		if root.right.payload != value {
			root.right.delete(value, decrement)
			return
		}
		if root.right.left == nil || root.right.right == nil {
			if root.right.left == nil {
				root.right = root.right.left
			} else {
				root.right = root.right.right
			}
			if decrement {
				root.size--
			}
			return
		}

		minVal := root.right.right.findMin().payload
		root.right.payload = minVal
		root.right.size--
		root.size--
		root.right.delete(minVal, false)

		return
	}

}

func (root *treeNode) findMin() *treeNode {
	if root.left == nil {
		return root
	}
	return root.left.findMin()
}

func (root *treeNode) getRandom() int {
	random := rand.Intn(root.size) + 1
	if random == root.size {
		return root.payload
	}

	if root.left == nil || random > root.left.size {
		return root.right.getRandom()
	}
	return root.left.getRandom()
}

func printNode(node *treeNode) {
	fmt.Printf("node value: %v, node size: %v, has left %v, has right %v \n", node.payload, node.size, node.left != nil, node.right != nil)
}

func main() {
	root := &treeNode{payload: 10, size: 1}
	root.insert(5)
	root.insert(15)
	root.insert(2)
	root.insert(7)
	//root.delete(5, true)
	printNode(root)
	printNode(root.left)
	printNode(root.right)
	printNode(root.left.left)
	//rintNode(root.left.right)

	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v, randomNodeValue: %v\n", i, root.getRandom())
	}
}
