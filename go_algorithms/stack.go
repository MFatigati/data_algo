/*

 */

package main

import "fmt"

type StackNode struct {
	Val  int
	Next *StackNode
}

type Stack struct {
	topNode *StackNode
	Size    int
}

func (s *Stack) push(newNode *StackNode) {
	if s.isEmpty() {
		s.topNode = newNode
		s.Size++
	} else {
		newNode.Next = s.topNode
		s.topNode = newNode
		s.Size++
	}
}

func (s *Stack) pop() *StackNode {
	toReturn := s.topNode
	s.topNode = s.topNode.Next
	return toReturn
}

func (s *Stack) isEmpty() bool {
	if s.Size == 0 {
		return true
	} else {
		return false
	}
}

func (s *Stack) peek() {
	fmt.Println(s.topNode.Val)
}

func main() {

	node1 := StackNode{}
	node1.Val = 1
	node2 := StackNode{}
	node2.Val = 2
	node3 := StackNode{}
	node3.Val = 3

	stack1 := Stack{}
	stack1.push(&node1)
	stack1.peek()
	stack1.push(&node2)
	stack1.peek()
	stack1.push(&node3)
	stack1.peek()
	fmt.Println(stack1.topNode.Next)
	fmt.Println(stack1.pop())
	stack1.peek()

}
