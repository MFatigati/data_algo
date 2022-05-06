package main

import "fmt"

type StringStack struct {
  data []string
}

func (s *StringStack) push(elem string) {
  s.data = append(s.data, elem)
}

func (s *StringStack) pop() string {
  popped := s.data[len(s.data) - 1]
  s.data = s.data[:len(s.data) - 1]
  return popped
}

func (s *StringStack) read() string {
  return s.data[len(s.data) - 1]
}

func reverseWord (str string) string {
  stack1 := StringStack{}
  reversed := ""
  for _, char := range str {
    stack1.push(string(char))
  }
  for len(stack1.data) > 0 {
    reversed += stack1.pop()
  }
  return reversed
}

func main () {

  fmt.Println(reverseWord("abcde"))

}