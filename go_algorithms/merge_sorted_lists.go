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

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	dummy := &ListNode{Val: -1}
// 	current := dummy
// 	l1current := list1
// 	l2current := list2

// 	for current != nil && l1current != nil && l2current != nil {
// 		if l1current.Val <= l2current.Val {
// 			current.Next = l1current
// 			current = current.Next
// 			l1current = l1current.Next
// 		} else if l2current.Val < l1current.Val {
// 			current.Next = l2current
// 			current = current.Next
// 			l2current = l2current.Next
// 		}
// 	}

// 	if l1current != nil {
// 		current.Next = l1current
// 	}

// 	if l2current != nil {
// 		current.Next = l2current
// 	}

// 	//printAllNodes(dummy.Next)
// 	return dummy.Next
// }

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := new(ListNode)
	current := dummy

	for list1 != nil || list2 != nil {
		fmt.Println(list1, list2)
		fmt.Println(current)
		if list1 == nil {
			current.Next = list2
			current = current.Next
			list2 = list2.Next
		} else if list2 == nil {
			current.Next = list1
			current = current.Next
			list1 = list1.Next
		}

		if list1.Val == list2.Val {
			fmt.Println("list1", list1)
			fmt.Println("here i am")
			current.Next = list1
			fmt.Println("current.next", current.Next)
			current.Next.Next = list2
			fmt.Println("current.Next.Next", current.Next.Next)
			current = current.Next.Next
			list1 = list1.Next
			fmt.Println("list1", list1)
			list2 = list2.Next
			fmt.Println("list2", list2)
			// I think yours is having trouble because on line 77 you assign list1 to `current.Next`,
			// then you set the `Next` value of `current.Next` to list2. So now list1.Next == list2.Next
		} else if list1.Val < list2.Val {
			current.Next = list1
			current = current.Next
			list1 = list1.Next
		} else if list2.Val < list1.Val {
			current.Next = list2
			current = current.Next
			list2 = list2.Next
		}
	}
	return current
}

func main() {

	node1 := ListNode{Val: 1}
	node2 := ListNode{Val: 2}
	node3 := ListNode{Val: 4}
	node4 := ListNode{Val: 1}
	node5 := ListNode{Val: 3}
	node6 := ListNode{Val: 4}

	node1.Next = &node2
	node2.Next = &node3

	node4.Next = &node5
	node5.Next = &node6

	mergeTwoLists(&node1, &node4)

}
