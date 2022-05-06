package main

import "fmt"

func sumRecursively (s []int) (sum int) {
  if len(s) == 1 {
    return s[0]
  } else {
    return s[0] + sumRecursively(s[1:])
  }
}

func reverseStrRecursively (s string) (reversed string) {
  if len(s) == 1 {
    return string(s[0])
  } 
  return string(s[len(s) - 1:]) + reverseStrRecursively(s[:len(s) - 1])
}

func countXsRecursively (s string) (total int) {
  if len(s) == 1 {
    if string(s[0]) == "x" {
      return 1
    } else {
      return 0
    }
  }

  if string(s[0]) == "x" {
    return 1 + countXsRecursively(s[1:])
  } else {
    return countXsRecursively(s[1:])
  }
}

func waysToClimb(numStairs int) (int) {
  if numStairs == 0 {
    return 0
  }
  if numStairs == 1 {
    return 1
  }
  if numStairs == 2 {
    return 2
  }
  if numStairs == 3 {
    return 4
  }
  return waysToClimb(numStairs - 1) + waysToClimb(numStairs - 2) + waysToClimb(numStairs - 3)
}


func totalChars(s []string) int {
  if len(s) == 1 {
    return len(s[0])
  }
  return len(s[0]) + totalChars(s[1:])
}

func onlyEvens (n []int) (result []int) {
  if len(n) == 1 {
    if n[0] % 2 == 0 {
      result = []int{n[0]}
      return result
    } else {
      return onlyEvens(n[1:])
    }    
  }

  if n[0] % 2 == 0 {
    result = append(onlyEvens(n[1:]), n[0])
    return result
  } else {
    return onlyEvens(n[1:])
  }
}


func triangularNums(triNum int) int {
  if triNum == 1 {
    return 1
  }
  if triNum == 2 {
    return 3
  }
  if triNum == 3 {
    return 6
  }
  
  return triNum + triangularNums(triNum - 1)
}

func firstX(s string, currentIdx int) (idx int) {
  if string(s[0]) == "x" {
    return currentIdx
  } else {
    currentIdx += 1
    return firstX(s[1:], currentIdx)
  }
}

func uniquePaths(rows int, columns int) int {
  if rows == 1 || columns == 1 {
    return 1
  }
  return uniquePaths(rows - 1, columns) + uniquePaths(rows, columns - 1)
}

func main() {

  // toSum := []int{1, 2, 3, 4, 5}
  // fmt.Println(sumRecursively(toSum))

  //fmt.Println(reverseStrRecursively("abcde"))
  //fmt.Println(countXsRecursively("axbxcxd"))
  //fmt.Println(waysToClimb(10))
  //fmt.Println(uniquePaths(10, 10))
  //fmt.Println(totalChars([]string{"ab", "c", "def", "ghij"}))
  // fmt.Println(onlyEvens([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
  // fmt.Println(triangularNums(7))
  fmt.Println(firstX("abcdefghijklmnopqrstuvwxyz", 0))

}