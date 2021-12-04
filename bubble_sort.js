function bubble_sort(arr) {
  let sorted = false;
  let lengthUntilSorted = arr.length - 1

  while (sorted === false) {
    sorted = true;
    for (let i = 0; i < lengthUntilSorted; i++) {
      if (arr[i] > arr[i + 1]) {
        // let lowerVal = arr[i + 1];
        // let higherVal = arr[i];
        // arr[i] = lowerVal;
        // arr[i + 1] = higherVal;
        [arr[i], arr[i + 1]] = [arr[i + 1], arr[i]];
        sorted = false;
      }
    }
    lengthUntilSorted--;
  }
  return arr;
}

console.log(bubble_sort([65, 55, 45, 35, 25, 15, 10]))