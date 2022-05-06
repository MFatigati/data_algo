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

func deleteDuplicates2(head *ListNode) *ListNode {

	forRemoval := map[int]bool{}

	if head == nil {
		return head
	}

	dummy := &ListNode{-1, nil}
	previous := dummy
	current := head
	previous.Next = current
	printAllNodes(dummy.Next)

	for current.Next != nil {
		_, exists := forRemoval[current.Val]
		if exists {
			current = current.Next
			previous.Next = current
		}
		if current.Val == current.Next.Val {
			forRemoval[current.Val] = true
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	_, exists := forRemoval[current.Val]
	if exists {
		previous.Next = nil
	}

	fmt.Println("transformed")
	printAllNodes(dummy.Next)
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
	// node1 := ListNode{Val: 1}
	// node2 := ListNode{Val: 1}
	// node3 := ListNode{Val: 1}
	// node4 := ListNode{Val: 2}
	// node5 := ListNode{Val: 3}
	// node1.Next = &node2
	// node2.Next = &node3
	// node3.Next = &node4
	// node4.Next = &node5

	// printAllNodes(&node1)

	deleteDuplicates2(&node1)

	// printAllNodes(&node1)

}
