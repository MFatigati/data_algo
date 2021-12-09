class Node {
  constructor(data) {
    this.data = data;
    this.nextNode = undefined;
    this.previousNode = undefined;
  }
}

let node_1 = new Node("once");
let node_2 = new Node("upon");
let node_3 = new Node("a");
let node_4 = new Node("time");
node_1.nextNode = node_2;
node_2.nextNode = node_3;
node_3.nextNode = node_4;
node_2.previousNode = node_1;
node_3.previousNode = node_2;
node_4.previousNode = node_3;

class LinkedList {
  constructor(firstNode) {
    this.firstNode = firstNode;
  }
  
  read(index) {
    let currentNode = this.firstNode;
    let currentIdx = 0;
    while (currentIdx < index) {
      currentNode = currentNode.nextNode;
      currentIdx += 1;
    }
    if (!currentNode) {
      return null;
    }
    return currentNode.data;
  }

  findIndexOfValue(value) {
    let currentNode = this.firstNode;
    let currentIdx = 0;
    while (currentNode.data !== value) {
      currentNode = currentNode.nextNode;
      currentIdx += 1;
    }
    if (!currentNode) {
      return null;
    }
    return currentIdx;
  }

  insertValueAtIndex(index, value) {
    let newNode = new Node(value);
    if (index === 0) {
      newNode.nextNode = this.firstNode;
      this.firstNode = newNode;
      return;
    }

    let currentNode = this.firstNode;
    let currentIdx = 0;
    while (currentIdx < index - 1) {
      currentNode = currentNode.nextNode;
      currentIdx++;
    }

    newNode.nextNode = currentNode.nextNode;
    currentNode.nextNode = newNode;
  }

  deleteNodeAtIndex(index) {
    if (index === 0) {
      this.firstNode = this.firstNode.nextNode
      return
    }

    let currentNode = this.firstNode;
    let currentIdx = 0;
    while (currentIdx < index - 1) {
      currentNode = currentNode.nextNode;
      currentIdx++;
    }
    currentNode.nextNode = currentNode.nextNode.nextNode;
  }

  printAllNodeValues() {
    let currentNode = this.firstNode;
    while (currentNode.data) {
      console.log(currentNode.data);
      if (!currentNode.nextNode) break;
      currentNode = currentNode.nextNode;
    }
  }

  returnLastValue() {
    let currentNode = this.firstNode;
    while (currentNode.nextNode) {
      currentNode = currentNode.nextNode;
      if (!currentNode.nextNode) {
        return currentNode.data;
      }
    }
  }

  /*
  while there is a next node
    set the first node to the next node
    set the current node's next node to the next node of the next node
  */
  reverseList() {
    let currentNode = this.firstNode;
    let priorFirst = this.firstNode;
    while (currentNode.nextNode) {
      let next = currentNode.nextNode;
      let nextNext = currentNode.nextNode.nextNode;
      this.firstNode = next;
      this.firstNode.nextNode = priorFirst;
      priorFirst = this.firstNode;
      currentNode.nextNode = nextNext;
    }
  }
}

class DoublyLinkedList {
  constructor(firstNode, lastNode) {
    this.firstNode = firstNode;
    this.lastNode = lastNode;
  }

  printInReverse() {
    let currentNode = this.lastNode
    while (true) {
      console.log(currentNode.data);
      if (!currentNode.previousNode) {
        break;
      }
      currentNode = currentNode.previousNode;
    }
  }
}

let list1 = new LinkedList(node_1);

// console.log(list1.read(3));
// console.log(list1.findIndexOfValue("time"));

list1.insertValueAtIndex(3, "purple");
// console.log(list1.findIndexOfValue("time"));
// console.log(list1.findIndexOfValue("purple"));
//list1.printAllNodeValues();

let list2 = new DoublyLinkedList(node_1, node_4);
// list2.printInReverse();

//console.log(list1.returnLastValue());
list1.reverseList();
list1.printAllNodeValues();