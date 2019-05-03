package main

import "fmt"

type ListNode struct {
	payload int
	next    *ListNode
}

func printList(head *ListNode) {
	for current := head; current != nil; current = current.next {
		fmt.Printf("%v -> ", current.payload)
	}
	fmt.Println()
}

func isPalindrom(head *ListNode) bool {
	lenght := 0
	var reservedHead *ListNode
	for current := head; current != nil; current = current.next {
		lenght++
		reservedHead = &ListNode{current.payload, reservedHead}
	}

	current := head
	reversedCurrent := reservedHead
	for i := 0; i < lenght/2; i++ {
		if current.payload != reversedCurrent.payload {
			return false
		}
		current = current.next
		reversedCurrent = reversedCurrent.next
	}
	return true
}

func main() {
	i5 := ListNode{payload: 1, next: nil}
	i4 := ListNode{payload: 2, next: &i5}
	i3 := ListNode{payload: 5, next: &i4}
	i2 := ListNode{payload: 2, next: &i3}
	i1 := ListNode{payload: 2, next: &i2}
	printList(&i1)

	fmt.Println(isPalindrom(&i1))
}
