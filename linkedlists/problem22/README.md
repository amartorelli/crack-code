# Problem
Implement an algorithm to find the nth to last element of a singly linked list

# Solution
The solution is to use two pointers, one that will scroll to the very end of the list (`tail`) counting how many elements there are in the list.
Once we know this information we can subtract the `nth` parameter from the total and see how many steps the second pointer must take.
Move then the pointer that many steps and return the node.