package main

import "fmt"

/*
Implement an algorithm to find the nth to last element of a singly linked list
*/

func (l *listElement) nthToLast(n int) *listElement {
	if l == nil {
		return nil
	}

	// tail will point to the end of the list
	tail := l
	// p will scroll through the list
	p := l

	// counter will be used to make sure the length of the list is at least n and to calculate what position the nth is
	counter := 0
	// send tail to the end of the list
	for tail.next != nil {
		tail = tail.next
		counter++
	}

	// if the list is smaller than n, return nil
	if counter < n {
		return nil
	}

	steps := counter - n
	for steps > 0 {
		p = p.next
		steps--
	}

	return p
}

func main() {
	nth := 3
	list := buildList(1, 2, 4, 5, 2, 10, 4)
	list.printList()
	fmt.Printf("%d to last is %d", nth, list.nthToLast(nth).val)
}
