class TreeNode {
  constructor(value, left=null, right=null) {
    this.value = value;
    this.leftChild = left;
    this.rightChild = right;
  }
}

function searchTree(searchValue, currentNode) {
  if (currentNode === null || currentNode.value === searchValue) {
    return currentNode
  } else if (searchValue < currentNode.value) {
    return searchTree(searchValue, currentNode.leftChild);
  } else {
    return searchTree(searchValue, currentNode.rightChild);
  }
}

function insertIntoTree(value, currentNode) {
  if (value < currentNode.value) {
    if (!currentNode.leftChild) {
      currentNode.leftChild = TreeNode(value);
    } else {
      insertIntoTree(value, currentNode.leftChild)
    }
  } else if (value > currentNode.value) {
    if (!currentNode.rightChild) {
      currentNode.rightChild = TreeNode(value);
    } else {
      insertIntoTree(value, currentNode.rightChild)
    }
  }
}

function findGreatest(currentNode) {
  if (currentNode.rightChild) {
    return findGreatest(currentNode.rightChild)
  } else {
    return currentNode.rightChild;
  }
}