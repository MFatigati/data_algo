let arr1 = [1, 2, 3, 4, 5, 6, 7, 8, 0, 3, 4, 5, 1];

function product3N3(arr) {
  let greatest = 0;
  for (let i = 0; i < arr.length; i++) {
    for (let j = i + 1; j < arr.length; j++) {
      for (let k = j + 1; k < arr.length; k++) {
        if ((arr[i] * arr[j] * arr[k]) > greatest) {
          console.log(arr[i], arr[j], arr[k])
          greatest = (arr[i] * arr[j] * arr[k]);
        }
      }
    }
  }
  return greatest;
}

// console.log(product3N3(arr1));

function product3N(arr) {
  arr = arr.sort();
  return arr[arr.length - 1] * arr[arr.length - 2] * arr[arr.length - 3];
}

console.log(product3N(arr1))

function findMissing(arr) {
  arr = arr.sort();
  for (let i = 0; i < arr.length; i++) {
    if (i === 0) {
      if (arr[i] !== arr[i + 1] - 1) {
        return arr[i] + 1;
      }
    } else {
      if (arr[i] !== arr[i - 1] + 1) {
        return arr[i - 1] + 1;
      }
    }
  }
}

//console.log(findMissing([9, 3, 2, 5, 6, 7, 1, 0, 4]))

function findGreatestN2(arr) {
  for (let i = 0; i < arr.length; i++) { // N
    let greatest = true;
    for (let j = i + 1; j < arr.length; j++) { // N
      if (arr[i] < arr[j]) {
        greatest = false;
      }
    }
    if (greatest === true) {
      return arr[i];
    }
  }
}

console.log(findGreatestN2([9, 3, 2, 5, 6, 7, 111, 0, 4]))

function findGreatestOlogN(arr) {
  arr = arr.sort(); // log N
  let greatest = arr[0];
  for (let i = 0; i < arr.length; i++) { // N
    if (arr[i] > greatest) {
      greatest = arr[i];
    }
  }
  return greatest;
}

console.log(findGreatestOlogN([9, 3, 2, 5, 6, 7, 111, 0, 4]))

function findGreatestN(arr) {
  let greatest = arr[0];
  for (let i = 0; i < arr.length; i++) { // N
    if (arr[i] > greatest) {
      greatest = arr[i];
    }
  }
  return greatest;
}

console.log(findGreatestN([9, 3, 2, 5, 6, 7, 111, 0, 4]))
