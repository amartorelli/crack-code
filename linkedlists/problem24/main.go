package main

import "fmt"

/*
You have two numbers represented by a linked list, where each node contains a sin- gle digit The digits are stored in reverse order, such that the 1’s digit is at the head of the list Write a function that adds the two numbers and returns the sum as a linked list
EXAMPLE
Input: (3 -> 1 -> 5) + (5 -> 9 -> 2)
Output: 8 -> 0 -> 8
*/

func sum(a, b *listElement) *listElement {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}

	var carryover int
	var r *listElement
	var result *listElement
	pa := a
	pb := b

	for pa != nil && pb != nil {
		res := pa.val + pb.val + carryover
		carryover = int(res / 10)
		v := &listElement{val: res % 10}

		// set head
		if result == nil {
			result = v
		}

		if r == nil {
			r = v
		} else {
			r.next = v
			r = r.next
		}

		pa = pa.next
		pb = pb.next
	}

	// finish appending the numbers from the list that still has elements
	for pa != nil {
		v := &listElement{val: pa.val}
		r.next = v
		pa = pa.next
	}
	for pb != nil {
		v := &listElement{val: pb.val}
		r.next = v
		pb = pb.next
	}

	// add carryover if exists
	if carryover > 0 {
		v := &listElement{val: carryover}
		r.next = v
	}

	return result
}

func main() {
	op1 := buildList(3, 1, 5)
	op2 := buildList(5, 9, 5)
	fmt.Println("first operator")
	op1.printList()
	fmt.Println("second operator")
	op2.printList()
	fmt.Println("result")
	sum(op1, op2).printList()
}
