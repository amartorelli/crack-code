# Problem
You have two numbers represented by a linked list, where each node contains a sin- gle digit The digits are stored in reverse order, such that the 1’s digit is at the head of the list Write a function that adds the two numbers and returns the sum as a linked list
EXAMPLE
Input: (3 -> 1 -> 5) + (5 -> 9 -> 2)
Output: 8 -> 0 -> 8

# Solution
The solution is to sum every element at the same position and keep track of the carryover which must be summed at the next iteration.
If any of the lists have extra elements, they get appended to the results at the end.
If there's any carryover left at the end, append that too.