let arr1 = [4, 2, 7, 1, 3];

/**
 * starting at index 1, iterate over all remaining index
 *  for each index
 *       store value at index in currentVal
 *       iterate backwards to zero
 *          if the backwards value is greater than currentVal, continue
 *          if the backwards value is less than the currentVal
 *            insert currentVal just to the right of this backwards val
 *            break out of this iteration of the loop
 * return modified array
 */


// function insertionSort(arr) {
//   for (let indexToInspect = 1; indexToInspect < arr.length; indexToInspect++) {
//     let currentVal = arr[indexToInspect];
//     let gapAt = indexToInspect;
//     for (let comparisonIdx = indexToInspect - 1; comparisonIdx >= 0; comparisonIdx--) {
//       if (arr[comparisonIdx] > currentVal) {
//         let valToShiftUp = arr[comparisonIdx];
//         let valToShiftDown = arr[gapAt];
//         arr[gapAt] = valToShiftUp;
//         arr[comparisonIdx] = valToShiftDown;
//         //[arr[comparisonIdx], arr[gapAt]] = [arr[gapAt], arr[comparisonIdx]];
//         gapAt--;
//       } else {
//         gapAt--;
//       }
//       console.log(arr);
//       }
//     }
//   return arr;
// }

function insertionSort(arr) {
  for (let indexToInspect = 1; indexToInspect < arr.length; indexToInspect++) {
    let currentVal = arr[indexToInspect];
    let comparisonIdx = indexToInspect - 1;
    while (comparisonIdx >= 0) {
      if (arr[comparisonIdx] > currentVal) {
        arr[comparisonIdx + 1] = arr[comparisonIdx];
        comparisonIdx--;
      } else {
        break;
      }
    }
    arr[comparisonIdx + 1] = currentVal;
  }
  return arr;
}

console.log(insertionSort(arr1))