package main

/*
Write code to remove duplicates from an unsorted linked list
FOLLOW UP
How would you solve this problem if a temporary buffer is not allowed?
*/

/*
	1 -> 2 -> 2 -> 0 -> nil
	A    A
	|    |
	p    |
	     n
*/

func (l *listElement) dedup() {
	// return straightaway if we have an empty list
	if l == nil {
		return
	}
	// return if we only have one element or we are analysing the last
	if l.next == nil {
		return
	}

	// hashmap to keep track of the ones we've already seen
	seen := make(map[int]struct{}, 0)

	// we add the first element to the list and start analysing from the second one
	seen[l.val] = struct{}{}

	// sets the previous node to the first element
	prev := l
	// sets the analysed node to the second one
	n := l.next

	// keep going until we reach the end
	for n != nil {
		// check if we've already seen the value
		_, found := seen[n.val]
		if found {
			// remove (jump) node
			prev.next = n.next
		} else {
			// otherwise we move the previous node pointer forward ready for the next iteration
			prev = prev.next
		}
		// adding the value to the seen map
		seen[n.val] = struct{}{}
		// in both cases we keep moving forward with the analysed node pointer
		n = n.next
	}
}

func main() {
	list := buildList(1, 2, 4, 5, 2, 10, 4)
	list.printList()
	list.dedup()
	list.printList()
}
