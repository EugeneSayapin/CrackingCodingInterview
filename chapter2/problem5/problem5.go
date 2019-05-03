package main

import "fmt"

type ListNode struct {
	Next  *ListNode
	Value int
}

func printList(head *ListNode) {
	for current := head; current != nil; current = current.Next {
		fmt.Printf("%v -> ", current.Value)
	}
}

func addReversed(head1 *ListNode, head2 *ListNode) *ListNode {
	carry := 0
	var resultHead, currentResult *ListNode
	for current1, current2 := head1, head2; current1 != nil || current2 != nil; {
		current1Digit, current2Digit := 0, 0

		if current1 != nil {
			current1Digit = current1.Value
		}

		if current2 != nil {
			current2Digit = current2.Value
		}

		sum := current1Digit + current2Digit + carry

		current := ListNode{nil, sum % 10}
		carry = sum / 10

		if currentResult != nil {
			currentResult.Next = &current
		}
		currentResult = &current

		if resultHead == nil {
			resultHead = currentResult
		}

		if current1 != nil {
			current1 = current1.Next
		}
		if current2 != nil {
			current2 = current2.Next
		}
	}
	if carry > 0 {
		lastNode := ListNode{nil, carry}
		currentResult.Next = &lastNode
	}
	return resultHead
}
func addLists(head1, head2 *ListNode) *ListNode {
	tail, carry := addListNodeHelper(head1, head2)
	if carry == 0 {
		return tail
	}
	first := ListNode {tail, carry}
	return &first
}

func addListNodeHelper(head1, head2 *ListNode) (*ListNode, int) {	
	if head1.Next == nil && head2.Next == nil {
		last := ListNode{nil, (head1.Value + head2.Value) % 10}
		return &last, (head2.Value + head2.Value)/10
	}
	tail, carry := addListNodeHelper(head1.Next, head2.Next)
	sum := head1.Value + head2.Value + carry;
	current := ListNode{tail, sum % 10}
	return &current,  sum / 10

}

func main() {
	// i13 := ListNode{nil, 6}
	// i12 := ListNode{&i13, 1}
	// i11 := ListNode{&i12, 7}
	// i23 := ListNode{nil, 2}
	// i22 := ListNode{&i23, 9}
	// //i21 := ListNode{&i22, 5}
	// printList(&i11)
	// fmt.Println("\n+")
	// printList(&i22)
	// fmt.Println("\n=")
	// printList(addReversed(&i11, &i22))

	i13 := ListNode{nil, 7}
	i12 := ListNode{&i13, 1}
	i11 := ListNode{&i12, 6}
	i23 := ListNode{nil, 5}
	i22 := ListNode{&i23, 9}
	i21 := ListNode{&i22, 2}
	printList(&i11)
	fmt.Println("\n+")
	printList(&i21)
	fmt.Println("\n=")
	printList(addReversed(&i11, &i21))
}
