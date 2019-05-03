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

func listLength(head *ListNode) int {
	length := 0
	for current := head; current != nil; current, length = current.next, length+1 {
	}
	return length
}

func findIntersection(head1, head2 *ListNode) *ListNode {
	length1 := listLength(head1)
	length2 := listLength(head2)

	var longListHead *ListNode
	var shortListHead *ListNode
	var nodesToSkip int

	if length1 > length2 {
		longListHead = head1
		shortListHead = head2
		nodesToSkip = length1 - length2
	} else {
		longListHead = head2
		shortListHead = head1
		nodesToSkip = length2 - length1
	}

	for i := 0; i < nodesToSkip; i++ {
		longListHead = longListHead.next
	}
	for current1, current2 := longListHead, shortListHead; current1 != nil; current1, current2 = current1.next, current2.next {
		if current1 == current2 {
			return current1
		}

	}
	return nil
}

func main() {

	i7 := ListNode{payload: 1, next: nil}
	i6 := ListNode{payload: 2, next: &i7}
	i5 := ListNode{payload: 7, next: &i6}
	i4 := ListNode{payload: 9, next: &i5}
	i3 := ListNode{payload: 5, next: &i4}
	i2 := ListNode{payload: 1, next: &i3}
	i1 := ListNode{payload: 3, next: &i2}

	ii2 := ListNode{payload: 6, next: &i5}
	//ii2 := ListNode{payload: 6, next: nil}
	ii1 := ListNode{payload: 4, next: &ii2}

	printList(&i1)
	printList(&ii1)

	intersection := findIntersection(&i1, &ii1)
	if intersection == nil {
		fmt.Println("No intersection")
		return
	}
	fmt.Printf("Intersection at element with payload = %v \n", intersection.payload)

}
