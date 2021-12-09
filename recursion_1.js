function sum(low, high) {
  if (low === high) return low;
  return high + sum(low, high - 1);
}

//console.log(sum(1, 10));

let array1 = [ 1, 2, 3,
  [4, 5, 6],
  7,
  [8,
    [9, 10, 11,
      [12, 13, 14]
    ] ],
  [15, 16, 17, 18, 19,
    [20, 21, 22,
      [23, 24, 25,
        [26, 27, 29]
      ], 30, 31 ], 32
  ], 33 ];

function recurseArrays(current) {
  if (typeof current === "number") {
    console.log(current);
  } else {
    for (let i = 0; i < current.length; i++) {
      recurseArrays(current[i]);
    }
  }
}

recurseArrays(array1);