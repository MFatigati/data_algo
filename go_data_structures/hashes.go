package main

import "fmt"

func findIntersection(slice1 []int, slice2 []int) (intersection []int) {
  lookup := map[int]bool{}
  for _, elem := range slice1 {
    lookup[elem] = true
  }
  for _, elem2 := range slice2 {
    if lookup[elem2] {
      intersection = append(intersection, elem2)
    }
  }
  return intersection
}

func findDup (strings []string) (dup string) {
  seen := map[string]bool{}

  for _, elem := range strings {
    if (seen[elem]) {
      return elem
    } else {
      seen[elem] = true
    }
  }
  return dup
}

func findMissingAlpha (string1 string) (missing string) {
  alphabet := "abcdefghijklmnopqrstuvwxyz"
  seen := map[string]bool{}
  for _, char := range string1 {
    seen[string(char)] = true
  }
  for _, alpha := range alphabet {
    if !seen[string(alpha)] {
      return string(alpha)
    }
  }
  return ""
}

func firstNonDup(string1 string) (nonDup string) {
  seen := map[string]int{}
  for _, elem := range string1 {
    if seen[string(elem)] >= 1 {
      seen[string(elem)] += 1
    } else {
      seen[string(elem)] = 1
    }
  }
  fmt.Println(seen)

  for _, elem := range string1 {
    if seen[string(elem)] == 1 {
      nonDup = string(elem)
      return nonDup
    }
  }
  return nonDup
}

func main() {
  // x := []int{1, 2, 3, 4, 5}
  // y := []int{0, 2, 4, 6, 8}
  // fmt.Println(findIntersection(x, y))

  // strings := []string{"a", "b", "c", "d", "c", "e", "f"}
  // fmt.Println(findDup(strings))

  //fmt.Println(findMissingAlpha("the quick brown box jumps over a lazy dog"))

  fmt.Println(firstNonDup("minimum"))
}

