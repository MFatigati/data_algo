package main

import (
	"fmt"
)

type mixed interface {
}

/*
COMMENT 2: It is important that the value we pass to `recurseArrays` not be a "mixed" interface itself, but rather a slice with "mixed" values, because you can range over slices, but not over interfaces. That is, you can range over slices of interfaces, but you can't range over interfaces directly.

For each value in the current slice, you use a type assertion to check whether its an int, or a slice full of "mixed" values. Again, you are not checking for "mixed" values themselves, but rather slice of "mixed" values. This is important, because only slices of "mixed" values, not "mixed" values themselves, can be passed to `recurseArrays`.

If the value is an int, print it. If it is a slice of "mixed" values, pass the slice.
*/

func recurseArrays(current []interface{}) {
  for _, currentElem := range current {
    switch current := currentElem.(type) {
      case int:
        fmt.Println(current)
      case []interface{}:
        recurseArrays(current)
    }
  }
}


func main() {

/*
COMMENT 1: To create the mixed multi-dimensional array, you need a type capable of storing integers AND slices. Normal slices only store values of a particular type. By creating an empty "mixed" interface type (defined above), we can then create a slice that stores anything that implements "mixed". So in the below, we are creating one slice, full of values that implement "mixed". Some of those values are are integers, some are slices, which are themselves full of "mixed" values.
*/

// multiDimensional := []mixed{ 1, 2, 3,
//   []mixed{4, 5, 6},
//   7,
//   []mixed{8,
//     []mixed{9, 10, 11,
//       []mixed{12, 13, 14},
//   } },
//   []mixed{15, 16, 17, 18, 19,
//     []mixed{20, 21, 22,
//       []mixed{23, 24, 25,
//         []mixed{26, 27, 29},
//       }, 30, 31 }, 32,
//   }, 33 };

multiDimensional := []interface{}{ 1, 2, 3,
  []interface{}{4, 5, 6},
  7,
  []interface{}{8,
    []interface{}{9, 10, 11,
      []interface{}{12, 13, 14},
  } },
  []interface{}{15, 16, 17, 18, 19,
    []interface{}{20, 21, 22,
      []interface{}{23, 24, 25,
        []interface{}{26, 27, 29},
      }, 30, 31 }, 32,
  }, 33 };
  
  
  recurseArrays(multiDimensional)

  slice1 := []int{1, 2, 3}
  something := make([]interface{}, len(slice1))
  fmt.Println(something)

}