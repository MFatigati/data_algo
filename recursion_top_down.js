let arr = [1, 2, 3, 4, 5];

function sumRecursively(arr) {
  if (arr.length === 1) {
    return arr[0]
  } else {
    return arr[0] + sumRecursively(arr.slice(1));
  }
}

// console.log(sumRecursively(arr))

let string1 = "abcde"

function reverseStrRecursively(str) {
  if (str.length === 1) {
    return str[0]
  } else {
    return reverseStrRecursively(str.slice(1)) + str[0];
  }
}

// console.log(reverseStrRecursively(string1));

let string2 = "axbxcxd"

function countXsRecursively(str) {
  // if (str.length === 1) {
  //   if (str[0] === "x") {
  //     return 1
  //   } else {
  //     return 0
  //   }  
  // }
  if (str.length === 0) return 0;

  if (str[0] === "x") {
    return 1 + countXsRecursively(str.slice(1));
  } else {
    return countXsRecursively(str.slice(1));
  }
}

// console.log(countXsRecursively(string2));

function waysToClimb(numStairs) {
  if (numStairs === 1) return 1;
  if (numStairs === 2) return 2;
  if (numStairs === 3) return 4;
  if (numStairs === 4) return 7;
  return waysToClimb(numStairs - 1) + waysToClimb(numStairs - 2) + waysToClimb(numStairs - 3);
}

//console.log(waysToClimb(10));

function getAnagramsRecursively(str) {

}

let arr2 = ["ab", "c", "def", "ghij"];

function totalChars(arr) {
  if (arr.length === 1) {
    return arr[0].length;
  }
  return arr[0].length + totalChars(arr.slice(1));
}

//console.log(totalChars(arr2));

let arr3 = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

function onlyEvens(arr) {
  if (arr.length === 1) {
    if (arr[0] % 2 === 0) {
      return [arr[0]];
    } else {
      return [];
    }    
  }

  if (arr[0] % 2 === 0) {
    let tempArr = onlyEvens(arr.slice(1));
    tempArr.unshift(arr[0]);
    return tempArr;
  } else {
    return onlyEvens(arr.slice(1));
  }
}

//console.log(onlyEvens(arr3))

function triangularNums(num) {
  if (num === 1) return 1;
  if (num === 2) return 3;
  if (num === 3) return 6;

  return num + triangularNums(num - 1);
}

//console.log(triangularNums(7));

function firstX(string, index = 0) {
  if (string[0] === "x") {
    return index;
  } else {
    return firstX(string.slice(1), index += 1);
  }
}

//console.log(firstX("abcdefghijklmnopqrstuvwxyz"))

function uniquePaths(rows, columns) {
  if (rows === 1 || columns === 1) {
    return 1;
  }
  return uniquePaths(rows - 1, columns) + uniquePaths(columns - 1, rows);
}

//console.log(uniquePaths(100, 2));


// can't get this...

function anagrams(str) {
 if (str.length === 1) return str[0];
 let collection = [];
 let anagramsOfSubstrings = anagrams(str.slice(1));
 
 for (let i = 0; i < anagramsOfSubstrings.length; i++) {
  let substr = anagramsOfSubstrings[i];

  for (let j = 0; j < substr.length; j++) {
    let newAnagram = substr.slice(0, j) + str[0] + substr.slice(j);
    collection.push(newAnagram);
  }
 }
 console.log(str);
 return collection;
}

console.log(anagrams('abcd'));