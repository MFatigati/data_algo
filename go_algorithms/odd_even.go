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

/*
input: a head node
output: a head node
rules: put all the odd node index nodes (starting at 1, 3, 5) first,
	followed by all the even index nodes (starting at 2, 4, 6...)

1 2 3 4 5
CurrentOdd
CurrentEven
NextOdd
1.Next = 2.Next
2.Next = 3.Next
3.Next = 2
CurrentOdd.Next = CurrentEven.Next
CurrentEven.Next = NextOdd.Next
NextOdd.Next = CurrentEven

1 3 2 4 5
3.Next = 4.Next
4.Next = 5.Next (nil)
5.Next = 3
1 3 5 2 4

Dummy1.Next
Dummy2.Next
LastOdd = dummy1
LastEven = dummy2
Current = head
Set counter to 1
iterate through list
	if the counter is odd
		LastOdd.Next = current
		LastOdd = LastOdd.Next
	if the counter is even
		LastEven.Next = current
		LastEven = LastEven.Next
	current = current.next
	counter++
1 2 3 4 5

Connect LastOdd.Next to dummy2.Next

*/

func oddEvenList(head *ListNode) *ListNode {

	dummy1 := &ListNode{Val: -1}
	dummy2 := &ListNode{Val: -2}

	lastOdd := dummy1
	lastEven := dummy2
	current := head
	counter := 1

	for current != nil {
		if counter%2 == 1 {
			lastOdd.Next = current
			lastOdd = lastOdd.Next
		}
		if counter%2 == 0 {
			lastEven.Next = current
			lastEven = lastEven.Next
		}
		current = current.Next
		counter++
	}
	lastEven.Next = nil
	lastOdd.Next = dummy2.Next
	return dummy1.Next
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

	printAllNodes(oddEvenList(&node1))

}
