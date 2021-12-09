package main

import "fmt"

type Node struct {
  data string
  nextNode *Node
}

type LinkedList struct {
  firstNode *Node
}

func (l *LinkedList) printAll() {
  var currentNode *Node = l.firstNode
  for true {
    fmt.Println(currentNode.data)
    if currentNode.nextNode == nil {
      break
    }
    currentNode = currentNode.nextNode
  }
}

func main () {
  
node_1 := Node{}
node_2 := Node{}
node_3 := Node{}
node_1.data = "test1"
node_1.nextNode = &node_2
node_2.data = "test2"
node_2.nextNode = &node_3
node_3.data = "test3"

list1 := LinkedList{&node_1}
list1.printAll()

}