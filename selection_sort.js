let arr1 = [4, 2, 7, 1, 3];

/*
set starting index to 0
set final index to length - 1
while starting index is lower than ending index
  store current starting index in temporary variable
  store value of starting index in temporary variable
  declare newLowerVal variable, set to value of starting index
  declare indexToSwap variable, set to value of starting index
  iterate over arr, from starting index to end
    if value at current index is less than value at starting index
      set newLowerVal to current index's value
      set indexToSwap to current index
  swap value at indexToSwap with value at starting index
  increment starting index
*/

function selectionSort(arr) {
  let startingIndex = 0;
  while (startingIndex < arr.length) {
    // let currentStartVal = arr[startingIndex];
    // // let newLowerVal = currentStartVal;
    let lowestNumIndex = startingIndex;
    for (let i = startingIndex; i < arr.length; i++) {
      if (arr[i] < arr[lowestNumIndex]) {
        lowestNumIndex = i;
        //indexToSwap = i;
      }
    }
    if (startingIndex !== lowestNumIndex) {
      [arr[startingIndex], arr[lowestNumIndex]] = [arr[lowestNumIndex], arr[startingIndex]]
    }
    startingIndex++
  }
  return arr;
}

console.log(selectionSort(arr1));