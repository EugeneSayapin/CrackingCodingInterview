package main

import "fmt"

type ListNode struct {
	payload int
	next    *ListNode
}

func findLoop(head *ListNode) *ListNode {
	slow, fast := head.next, head.next.next
	for ; fast.next != nil && fast.next.next != nil; slow, fast = slow.next, fast.next.next {
		if slow == fast {
			break
		}
	}
	if fast.next == nil || fast.next.next == nil {
		return nil
	}

	for slow = head; ; slow, fast = slow.next, fast.next {
		if slow == fast {
			return slow
		}
	}

}

func main() {
	i1 := ListNode{payload: 1, next: nil}
	i2 := ListNode{payload: 2, next: nil}
	i3 := ListNode{payload: 3, next: nil}
	i4 := ListNode{payload: 4, next: nil}
	i5 := ListNode{payload: 5, next: nil}
	i6 := ListNode{payload: 6, next: nil}
	i7 := ListNode{payload: 7, next: nil}
	i8 := ListNode{payload: 8, next: nil}
	i9 := ListNode{payload: 9, next: nil}

	i1.next = &i2
	i2.next = &i3
	i3.next = &i4
	i4.next = &i5
	i5.next = &i6
	i6.next = &i7
	i7.next = &i8
	i8.next = &i9
	i9.next = nil

	loopStart := findLoop(&i1)

	if loopStart == nil {
		fmt.Println("No loop")
		return
	}

	fmt.Printf("Loop at position: %v", loopStart.payload)
}
