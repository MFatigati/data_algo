package main

import "fmt"

type TreeNode struct {
  value int
  leftChild *TreeNode
  rightChild *TreeNode
}



func insertIntoTree(value int, currentNode *TreeNode) {
  if value < currentNode.value {
    if currentNode.leftChild == nil {
      newNode := TreeNode{}
      newNode.value = value
      currentNode.leftChild = &newNode
    } else {
      insertIntoTree(value, currentNode.leftChild)
    }
  }
  if value > currentNode.value {
    if currentNode.rightChild == nil {
      newNode := TreeNode{}
      newNode.value = value
      currentNode.rightChild = &newNode
    } else {
      insertIntoTree(value, currentNode.rightChild)
    }
  }
}

func main() {

node1 := TreeNode{}
node2 := TreeNode{}
node1.value = 50
node2.value = 75
rootNode := TreeNode{25, &node1, &node2}

fmt.Println(rootNode);

node3 := TreeNode{}
node3.value = 60
insertIntoTree(60, &rootNode)
fmt.Println(rootNode.rightChild.leftChild.value)

}