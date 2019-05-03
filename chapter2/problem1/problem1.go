package main

import "fmt"

type ListNode struct {
	Next  *ListNode
	Value int
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Value)
		head = head.Next
	}
}

func removeDuplicates(head *ListNode) {
	for i := head; i != nil; {
		jPrev := i
		for j := i.Next; j != nil; {
			if i.Value == j.Value {
				jPrev.Next = j.Next
			} else {
				jPrev = j
			}
			j = j.Next
		}
		i = i.Next
	}
}

func main() {
	i5 := ListNode{nil, 5}
	i4 := ListNode{&i5, 4}
	i3 := ListNode{&i4, 3}
	i2 := ListNode{&i3, 5}
	i1 := ListNode{&i2, 1}
	printList(&i1)
	removeDuplicates(&i1)
	printList(&i1)
}
