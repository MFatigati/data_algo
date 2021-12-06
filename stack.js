class Linter {
  constructor(str) {
    this.values = [];
    this.stack = [];
    this.openingBraces = ["{", "[", "("];
    this.closingBraces = ["}", "]", ")"];
  }

  lint(str) {
    this.values = str.split("");
    for (let i = 0; i < this.values.length; i++) {
      let char = this.values[i];
      if (this.irrelevantValue(char)) {
        continue;
      }
      if (this.openingBraces.includes(char)) {
        this.stack.push(char)
      } else if (this.closingBraces.includes(char)) {
        let comparison = this.stack.pop();
        if (!comparison) {
          return `Error 2: closing brace ${char} never preceded by corresponding opening brace.`
        } else if (this.openingBraces.indexOf(comparison) !== this.closingBraces.indexOf(char)) {
          return `Error 3: closing brace ${char} not immediately preceded by corresponding opening brace.`
        } 
      }
      console.log(this.stack)
    }
    if (this.stack.length) {
      return `Error 1: opening brace ${this.stack[this.stack.length - 1]} with no closing brace.`
    }
    console.log(this.stack);
    return "No errors."
  }

  irrelevantValue(char) {
    return !this.openingBraces.includes(char) && !this.closingBraces.includes(char);
  }
}

let linter1 = new Linter();
console.log(linter1.lint("(var x = {y: [1, 2, 3]})"));