/*
P
	I:
	O:
E
D
A
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {

	dummy := &ListNode{Val: 1}
	current := dummy
	current.Next = head

	for current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return dummy.Next
}

func main() {

	node1 := ListNode{Val: 1}
	node2 := ListNode{Val: 2}
	node3 := ListNode{Val: 6}
	node4 := ListNode{Val: 3}
	node5 := ListNode{Val: 4}
	node6 := ListNode{Val: 5}
	node7 := ListNode{Val: 6}
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4
	node4.Next = &node5
	node5.Next = &node6
	node6.Next = &node7

	removeElements(&node1, 6)

}
