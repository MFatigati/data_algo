package main

import (
	"fmt"
	"sort"
)

func productThreeGreatest(slice1 []int) (product int) {
  sort.Ints(slice1)
  return slice1[len(slice1) - 1] * slice1[len(slice1) - 2] * slice1[len(slice1) - 3]
}

func findMissing(slice1 []int) int {
  sort.Ints(slice1)
  for i, elem := range slice1 {
    if i == 0 {
      if elem != slice1[i + 1] - 1 {
        return slice1[i + 1] - 1
      }
    } else {
      if elem != slice1[i - 1] + 1 {
        return slice1[i - 1] + 1
      }
    }
  }
  return 0
}

func findGreatestN2(slice1 []int) (greatest int) {
  for i := 0; i < len(slice1); i++ {
    isGreatest := true
    for j := 0; j < len(slice1); j++ {
      if slice1[i] < slice1[j] {
        isGreatest = false
      }
    }
    if isGreatest == true {
      greatest = slice1[i]
    } 
  }
  return greatest
}

func findGreatestNlogN(slice1 []int) (greatest int) {
  sort.Ints(slice1)
  greatest = slice1[len(slice1) - 1]
  return greatest
}

func findGreatestN(slice1 []int) (greatest int) {
  greatest = 0
  for _, elem := range slice1 {
    if elem > greatest {
      greatest = elem
    }
  }
  return greatest
}

func main() {
  // fmt.Println(productThreeGreatest([]int{1, 2, 3, 4, 5, 6, 7, 8, 0, 3, 4, 5, 1}))
  // fmt.Println(findMissing([]int{9, 3, 2, 5, 6, 7, 1, 0, 4}))
  fmt.Println(findGreatestN2([]int{9, 3, 2, 5, 6, 7, 1, 0, 4}))
  fmt.Println(findGreatestNlogN([]int{9, 3, 2, 5, 6, 7, 1, 0, 4}))
  fmt.Println(findGreatestN([]int{9, 3, 2, 5, 6, 7, 1, 0, 4}))

}
