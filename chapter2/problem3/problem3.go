package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

func printList(head *ListNode) {
	for current := head; current != nil; current = current.Next {
		fmt.Println(current.Value)
	}
}

func deleteMiddle(head *ListNode) {
	middle := head
	shouldMove := false
	var oneBeforeMiddle *ListNode = nil
	for current := head; current != nil; current = current.Next {
		if shouldMove {
			oneBeforeMiddle = middle
			middle = middle.Next
			shouldMove = false
		} else {
			shouldMove = true
		}
	}
	oneBeforeMiddle.Next = middle.Next

}

func main() {
	i9 := ListNode{9, nil}
	i8 := ListNode{8, &i9}
	i7 := ListNode{7, &i8}
	i6 := ListNode{6, &i7}
	i5 := ListNode{5, &i6}
	i4 := ListNode{4, &i5}
	i3 := ListNode{3, &i4}
	i2 := ListNode{2, &i3}
	i1 := ListNode{1, &i2}
	printList(&i1)
	deleteMiddle(&i1)
	printList(&i1)

}
