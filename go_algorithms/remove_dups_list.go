/*
P
	I:
	O:
E
D
A

Set current to head
set next to head.next

Until current.next == nil
	if current.next.value equals current.value
		set current.next to current.next.next
	else current = current.next



*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// While the next node is not nil
// if the value of the next node is the same as the current node
// skip over the next node by making the next value of the current node the next value of the next node
// otherwise, just advance to the next node by changing the current node to the next node

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return head
	}

	current := head

	for current.Next != nil {
		if current.Next.Val == current.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head

}

func main() {

	node1 := ListNode{Val: 1}
	node2 := ListNode{Val: 1}
	node3 := ListNode{Val: 2}
	node4 := ListNode{Val: 3}
	node5 := ListNode{Val: 3}
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	deleteDuplicates(&node1)

	fmt.Println(node1)
}
