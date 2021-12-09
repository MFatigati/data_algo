package main

import "fmt"

func fibDynamic (fibNum int, memo map[int]int) int {
  fmt.Println("RECURSION")
  if fibNum == 0 || fibNum == 1 {
    return fibNum
  } else {
    _, ok := memo[fibNum]
    if !ok {
      memo[fibNum] = fibDynamic(fibNum - 1, memo) + fibDynamic(fibNum - 2, memo)
    }
    return memo[fibNum]
  }
}

func golomb(n int, memo map[int]int) int {
  if n == 1 {
    return 1
  } else {
    _, ok := memo[n]
    if !ok {
      memo[n] = 1 + golomb(n - golomb(golomb(n - 1, memo), memo), memo);
    }
    return memo[n]
  }
}

func uniquePaths(rows int, columns int, memo map[string]int) int {
  if rows == 1 || columns == 1 {
    return 1
  }
  _, ok := memo[fmt.Sprintf("%d, %d", rows, columns)]
  if !ok {
    memo[fmt.Sprintf("%d, %d", rows, columns)] = uniquePaths(rows - 1, columns, memo) + uniquePaths(rows, columns - 1, memo)
  } 
  return memo[fmt.Sprintf("%d, %d", rows, columns)]
}

func main () {
  // fmt.Println(fibDynamic(5, map[int]int{}))
  // fmt.Println(fibDynamic(7, map[int]int{}))
  // fmt.Println(golomb(5, map[int]int{}))
  fmt.Println(uniquePaths(100, 2, map[string]int{}))
}