/*
P
	I:
	O:
E
D
A

set currentHead to head
set current to head
set previous to nil
set next to current.next

while current.next does not equal nill
	set current.next to currentHead
	set currentHead to current.next

*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var previous *ListNode
	current := head
	next := head.Next

	for current.Next != nil {
		current.Next = previous
		previous = current
		current = next
		next = next.Next
	}

	current.Next = previous

	return current

}

func main() {

	node1 := ListNode{Val: 1}
	node2 := ListNode{Val: 2}
	node3 := ListNode{Val: 3}
	node4 := ListNode{Val: 4}
	node5 := ListNode{Val: 5}
	node6 := ListNode{Val: 6}
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5
	node5.Next = &node6

	fmt.Println(reverseList(&node1))

}
