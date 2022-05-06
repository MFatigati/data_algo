/*

while count > 0
	traverse the list
		when node.Next.Next = nil
		set newHead to node.Next
		set node.Next to nil
		set newHead.Next to head
		decrement count

*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil || k == 0 {
		return head
	}

	lenList := 0

	currentNode := head

	// getting the length of the current list

	for currentNode != nil {
		currentNode = currentNode.Next
		lenList++
	}

	// getting the number of necessary rotations for when k > length of the list
	if k > lenList {
		k = k % lenList
	}

	count := k

	var newHead *ListNode = head
	var oldHead *ListNode = head

	// until you have done the necessary number of rotations
	for count > 0 {

		currentNode := oldHead

		for currentNode.Next.Next != nil {
			currentNode = currentNode.Next
		}

		// put tail at the beginning of the list
		newHead = currentNode.Next
		currentNode.Next = nil
		newHead.Next = oldHead
		oldHead = newHead

		count--
	}

	return newHead
}

func printAllNodes1(head *ListNode) {
	current := head
	for current != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
}

func main() {

	node1 := ListNode{Val: 1}
	node2 := ListNode{Val: 2}
	node3 := ListNode{Val: 3}
	node4 := ListNode{Val: 4}
	node5 := ListNode{Val: 5}
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5

	printAllNodes1(rotateRight(&node1, 6))

}
