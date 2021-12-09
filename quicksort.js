function partition(arr, leftPointer, rightPointer) {
  let pivotIdx = rightPointer;
  let pivotValue = arr[pivotIdx];
  rightPointer--;

  while (true) {
    while (arr[leftPointer] < pivotValue) {
      leftPointer += 1;
    }
    while (arr[rightPointer] > pivotValue) {
      rightPointer -= 1;
    }

    if (leftPointer >= rightPointer) {
      // if the left pointer has gone beyond or equal the right pointer
      // we don't need to swap any values, and we can stop the loop because
      // everything to the right of the right pointer is greater than anything
      // to the left of the left pointer, as they were both working "inwards" towards the
      // same value, just with different conditions
      break
      /*
      If both pointers have stopped and not met, we know that the right pointer has stopped
      at something less than what the left pointer has stopped at, because the right pointer
      was going until it reached something less than the partition value, and the left pointer
      was going until it reached something greater than the partition value. If the pointers
      did not meet, it means something greater is to the left of something lesser, so those
      values need to be swapped
       */
    } else {
      [arr[leftPointer], arr[rightPointer]] = [arr[rightPointer], arr[leftPointer]];
      // we also need to advance our left pointer, and continue on in the loop, until
      // we reach an index where the pointers do meet or overlap
      leftPointer += 1;
    };
  }
  [arr[leftPointer], arr[pivotIdx]] = [arr[pivotIdx], arr[leftPointer]];
  return leftPointer;
}

let arr1 = [0, 5, 2, 1, 6, 3, 2, 5, 2, 1, 9, 6];

function quicksort(arr, leftPointer, rightPointer) {
  if (rightPointer - leftPointer <= 0) {
    return
  }

  pivotIdx = partition(arr, leftPointer, rightPointer);

  quicksort(arr, leftPointer, pivotIdx - 1);
  quicksort(arr, pivotIdx + 1, rightPointer);
}

quicksort(arr1, 0, arr1.length - 1);

console.log(arr1);