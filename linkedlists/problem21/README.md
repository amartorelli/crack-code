# Problem
Write code to remove duplicates from an unsorted linked list
FOLLOW UP
How would you solve this problem if a temporary buffer is not allowed?

# Solution
The solution solves directly the follow up question, thus avoids the usage of an extra buffer. It keeps track of the previous node instead at each iteration so that the node can be jumped if already seen.