class Stack {
  constructor() {
    this.data = [];
  }

  push(elem) {
    this.data.push(elem);
  }

  pop() {
    return this.data.pop();
  }

  read() {
    this.data[this.data.length - 1];
  }
}

function reverseWord(str) {
  let stack = new Stack();
  let reversed = "";
  for (let i = 0; i < str.length; i++) {
    stack.push(str[i]);
  }
  for (let i = stack.data.length - 1; i >= 0; i--) {
    reversed += stack.pop();
  }
  return reversed;
}

console.log(reverseWord("abcde"));