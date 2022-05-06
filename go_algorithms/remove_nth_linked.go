/*
P
	I:
	O:
E
D
A
*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printAllNodes(head *ListNode) {
	current := head
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if head.Next == nil {
		return nil
	}

	dummy := &ListNode{-1, head}
	dummy.Next = head

	if head.Next == nil {
		return nil
	}

	length := 0
	seen := map[int]*ListNode{}
	current := head

	for current != nil {
		length++
		seen[length] = current
		current = current.Next
	}

	nodeToRemove := length - n + 1

	if n == 1 {
		seen[nodeToRemove-1].Next = nil
	} else if length == n {
		dummy.Next = seen[2]
	} else {
		seen[nodeToRemove-1].Next = seen[nodeToRemove+1]
	}

	return dummy.Next
}

func main() {

	node1 := ListNode{Val: 1}
	node2 := ListNode{Val: 2}
	node3 := ListNode{Val: 3}
	node4 := ListNode{Val: 3}
	node5 := ListNode{Val: 4}
	node6 := ListNode{Val: 5}
	node7 := ListNode{Val: 5}
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5
	node5.Next = &node6
	node6.Next = &node7
	printAllNodes(&node1)

	removeNthFromEnd(&node1, 2)
	printAllNodes(&node1)

}
