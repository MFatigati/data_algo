function fib(num) {
  if (num === 1 || num === 2) {
    return 1
  } else {
    return fib(num - 1) + fib(num - 2)
  }
}

// console.log(fib(7));

function fibDynamic(num, memo) {
  console.log("recursion");
  if (num === 1 || num === 2) {
    return 1
  } else if (num in memo) {
    return memo[num]
  } else {
    memo[num] = fibDynamic(num - 1, memo) + fibDynamic(num - 2, memo)
    return memo[num]
  }
}

// console.log(fibDynamic(7, {}));

// EXERCISE 1
// store `add_until_100(array[1, array.length - 1])` in a variable before the first
// if statement

function golomb(n, memo) {
  if (n === 1) {
    return 1
  } else {
    if (!memo[n]) {
      memo[n] = 1 + golomb(n - golomb(golomb(n - 1, memo), memo), memo);
    }
    return memo[n];
  }
}

//console.log(golomb(5, {}))

function uniquePaths(rows, columns, memo) {
  if (rows === 1 || columns === 1) {
    return 1;
  }
  if (!memo[`${rows - 1}/${columns}`] && !memo[`${rows}/${columns - 1}`]) {
    memo[`${rows - 1}/${columns}`] = uniquePaths(rows - 1, columns, memo)
    memo[`${rows}/${columns - 1}`] = uniquePaths(columns - 1, rows, memo)
  }
  return memo[`${rows - 1}/${columns}`] + memo[`${rows}/${columns - 1}`]
}

console.log(uniquePaths(100, 2, {}))
