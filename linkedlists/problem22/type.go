package main

import "fmt"

// listElement represents an element in the list
type listElement struct {
	val  int
	next *listElement
}

func (l *listElement) printList() {
	if l == nil {
		fmt.Printf("nil\n")
		return
	}
	fmt.Printf("%d -> ", l.val)
	l.next.printList()
}

// buildList builds a list of listElement starting from a list of integers
func buildList(nn ...int) *listElement {
	var head *listElement
	var previous *listElement

	for _, n := range nn {
		l := &listElement{val: n}
		if previous != nil {
			previous.next = l
		}
		if head == nil {
			head = l
		}
		previous = l
	}
	return head
}
