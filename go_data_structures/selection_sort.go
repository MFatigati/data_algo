package main

import "fmt"

func selectionSort(sliceToSort []int) []int {
  startingI := 0
  for startingI < len(sliceToSort) - 1 {
    startingVal := sliceToSort[startingI]
    newLowestVal := startingVal
    indexToSwap := startingI
    for i := startingI; i < len(sliceToSort); i++ {
      if sliceToSort[i] < newLowestVal {
        newLowestVal = sliceToSort[i]
        indexToSwap = i
      }
    }
    if newLowestVal != startingVal {
      sliceToSort[startingI], sliceToSort[indexToSwap] = newLowestVal, startingVal
    }
    startingI++
  }
  return sliceToSort
}

func main() {

slice1 := []int{4, 2, 7, 1, 3}
fmt.Println(selectionSort(slice1))


}