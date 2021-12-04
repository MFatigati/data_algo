function binarySearch(arr, searchVal) {
  console.log(arr)
  let steps = 0;
  let lowerBound = 0;
  let upperBound = arr.length - 1;

  while (lowerBound <= upperBound) {
    steps++;
    let midpoint = Math.round((upperBound + lowerBound) / 2)
    let valAtMidpoint = arr[midpoint]
    console.log(steps, midpoint);
    if (searchVal === valAtMidpoint) {
      return midpoint
    } else if (searchVal < valAtMidpoint) {
      upperBound = midpoint - 1
    } else if (searchVal > valAtMidpoint) {
      lowerBound = midpoint + 1
    }
  }
  return undefined;
}

console.log(binarySearch([...Array(100000)].map((elem, i, arr) => arr[i] = i + 1), 0))