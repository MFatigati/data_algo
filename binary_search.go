package main

import "fmt"

func binarySearch(slice1 []int, desired int) (val int, exists bool) {
  fmt.Println(slice1)
  var lowerBound int = 0
  var upperBound int = len(slice1) - 1

  for lowerBound <= upperBound {
    var midpoint int = (upperBound + lowerBound) / 2
    valAtMid := slice1[midpoint]
    if valAtMid == desired {
      return midpoint, true
    } else if valAtMid < desired {
      fmt.Println("value too low")
      lowerBound = midpoint + 1
    } else if valAtMid > desired {
      fmt.Println("value too low")
      upperBound = midpoint - 1
    }
  }
  return 0, false
}

func main () {
  testSlice := make([]int, 100)

  for i, _ := range testSlice {
    testSlice[i] = i + 1
  }

  fmt.Println(binarySearch(testSlice, 200))
}