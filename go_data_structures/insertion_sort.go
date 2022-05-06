package main

import "fmt"

func insertionSort(slice []int) []int {
  for i := 1; i < len(slice); i++ {
    currentVal := slice[i]
    comparisonIdx := i - 1
    for comparisonIdx >= 0 {
      if slice[comparisonIdx] > currentVal {
        slice[comparisonIdx + 1] = slice[comparisonIdx]
        comparisonIdx--
      } else {
        break;
      }
    }
    slice[comparisonIdx + 1] = currentVal
  }
  return slice
}

func main () {
  slice1 := []int{4, 2, 7, 1, 3}
  fmt.Println(insertionSort(slice1))
}