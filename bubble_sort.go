package main

import "fmt"

func bubble_sort(arr []int) (sortedArr []int) {
  var sorted bool = false
  var lengthUntilSorted = len(arr) - 1

  for sorted == false {
    sorted = true
    for i := 0; i < lengthUntilSorted; i++ {
      if arr[i] > arr[i + 1] {
        arr[i], arr[i + 1] = arr[i + 1], arr[i]
        sorted = false
      }
    }
    lengthUntilSorted--
  }
  sortedArr = arr[:]
  return sortedArr
}

func main () {

var arr1 [7]int = [7]int{10, 15, 25, 35, 45, 55, 65}
fmt.Println(bubble_sort(arr1[0:]))

}