package main

import (
	"fmt"
	"strings"
)

type Linter struct {
  values []string
  stack []string
  openingBraces map[string]bool
  closingBraces map[string]bool
}

func (linter *Linter) irrelevantValue(char string) bool {
  return !linter.openingBraces[char] && !linter.closingBraces[char]
}

func (linter *Linter) lint() (result string) {
  for _, char := range linter.values {
    if linter.irrelevantValue(char) {
      continue
    } else if linter.openingBraces[char] {
      linter.stack = append(linter.stack, char)
    } else if linter.closingBraces[char] {
      if (len(linter.stack) <= 0) {
        return fmt.Sprintf("Error 2: closing brace %v never preceded by corresponding opening brace.", char)
      } else {
        poppedComparison := linter.stack[len(linter.stack) - 1]
        linter.stack = linter.stack[0:len(linter.stack) - 1]
        fmt.Println(poppedComparison, char)
        if (char == "}") {
          if (poppedComparison != "{") {
            return `Error 3: closing brace } not immediately preceded by corresponding opening brace.`
          }
        }
        if (char == "]") {
          if (poppedComparison != "[") {
            return `Error 3: closing brace ] not immediately preceded by corresponding opening brace.`
          }
        }
        if (char == ")") {
          if (poppedComparison != "(") {
            return `Error 3: closing brace ] not immediately preceded by corresponding opening brace.`
          }
        }
      }
    }
  }
  if len(linter.stack) > 0 {
    return fmt.Sprintf("Error 1: opening brace %v with no closing brace.", linter.stack[0])
  }
  return "No errors"
}

func makeLinter(toInspect string) Linter {
  linter1 := Linter{}
  linter1.values = strings.Split(toInspect, "")
  linter1.openingBraces = map[string]bool{"{": true, "[": true, "(": true}
  linter1.closingBraces = map[string]bool{"}": true, "]": true, ")": true}
  return linter1
}


func main() {
  var linter1 Linter = makeLinter("(var x = {y: [1, 2, 3]})}")
  fmt.Println(linter1.lint())
}