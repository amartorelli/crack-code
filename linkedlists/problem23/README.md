# Problem
Implement an algorithm to delete a node in the middle of a single linked list, given only access to that node
EXAMPLE
Input: the node ‘c’ from the linked list a->b->c->d->e
Result: nothing is returned, but the new linked list looks like a->b->d->e

# Solution
I spent few minutes thinking of an approach and I couldn't find any because the biggest problem is that there's no way you can know the previous section of the list since the only entry point is the node you need to delete.
I looked at the solution for this one and then implemented it. After looking, the solution was obvious, but when you're focusing a lot trying to see if there's any property of linked lists you could take advantage of, you don't think at such a simple way to resolve this problem.
The solution is to copy the next node's content to the entry point and then bypass completely the next node (point the entry point's next to its next.next).