package main

import "fmt"

func reverseSlice(slice1 []int) []int {
  startIdx := 0
  endIdx := len(slice1) - 1
  for startIdx < endIdx {
    slice1[endIdx], slice1[startIdx] = slice1[startIdx], slice1[endIdx]
    startIdx++
    endIdx--
  }
  return slice1
}

func main() {
  fmt.Println(reverseSlice([]int{1, 2, 3, 4, 5}))
}