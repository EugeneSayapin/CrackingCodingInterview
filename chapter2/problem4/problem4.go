package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

func printList(head *ListNode) {
	for current := head; current != nil; current = current.Next {
		fmt.Printf("%v -> ", current.Value)
	}
	fmt.Println()
}

func partition(head *ListNode, elem int) {
	divider := head
	for current := head; current != nil; current = current.Next {
		if current.Value < elem {
			temp := divider.Value
			divider.Value = current.Value
			current.Value = temp
			divider = divider.Next
		}
	}
}

func main() {
	i1 := ListNode{3, nil}
	i2 := ListNode{5, nil}
	i3 := ListNode{8, nil}
	i4 := ListNode{5, nil}
	i5 := ListNode{10, nil}
	i6 := ListNode{2, nil}
	i7 := ListNode{1, nil}
	i1.Next = &i2
	i2.Next = &i3
	i3.Next = &i4
	i4.Next = &i5
	i5.Next = &i6
	i6.Next = &i7
	printList(&i1)
	partition(&i1, 5)
	printList(&i1)

}
