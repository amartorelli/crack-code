package main

import "fmt"

/*
Implement an algorithm to delete a node in the middle of a single linked list, given only access to that node
EXAMPLE
Input: the node ‘c’ from the linked list a->b->c->d->e
Result: nothing is returned, but the new linked list looks like a->b->d->e
*/

// delNode deletes the node passed as parameter
func delNode(n *listElement) {
	// this can't be achieved if the node is nil or if it doesn't have a next element
	if n == nil {
		return
	}
	if n.next == nil {
		return
	}

	// copy the next node and delete it
	n.val = n.next.val
	n.next = n.next.next
}

// nth returns the nth element in the list
func (l *listElement) nth(n int) *listElement {
	if l == nil {
		return nil
	}

	p := l
	for n-1 > 0 {
		p = p.next
		n--
	}
	return p
}

func main() {
	nth := 3
	list := buildList(1, 2, 4, 5, 2, 10, 4)
	list.printList()

	// get a node from the list
	param := list.nth(nth)
	fmt.Printf("input: %dth element = %d\n", nth, param.val)
	delNode(param)
	fmt.Println("result:")
	list.printList()
}
