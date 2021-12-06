function findIntersection(arr1, arr2) {
  let lookup = {};
  let intersection = [];
  arr1.forEach(elem => lookup[elem] = true);
  arr2.forEach(elem => {
    if (lookup[elem]) {
      intersection.push(elem);
    }
  })

  return intersection;
}

// console.log(findIntersection([1, 2, 3, 4, 5], [0, 2, 4, 6, 8]));

function findDup(arr1) {
  let seen = {};

  for (let i = 0; i < arr1.length; i++) {
    let currentVal = arr1[i];
    if (seen[currentVal]) {
      return currentVal
    } else {
      seen[currentVal] = true;
    }
  }

  return false;
}

//console.log(findDup(["a", "b", "c", "d", "c", "e", "f"]));

var a = 97;
var charArray = {};
for (var i = 0; i<26; i++)
    charArray[String.fromCharCode(a + i)] = false;
//console.log(charArray);

function findMissingAlpha(string) {
  string = string.toLowerCase();
  for (let i = 0; i < string.length; i++) {
    charArray[string[i]] = true;
  }
  for (let key in charArray) {
    if (charArray[key] === false)
    return key
  }
}

//console.log(findMissingAlpha("the quick brown box jumps over a lazy dog"))

function firstNonDUp(string) {
  let seen = {};
  for (let i = 0; i < string.length; i++) {
    seen[string[i]] = seen[string[i]] ? seen[string[i]] += 1 : 1;
  }
  console.log(seen)
  for (let i = 0; i < string.length; i++) {
    if (seen[string[i]] === 1) {
      return string[i]
    }
  }
}

// console.log(firstNonDUp("minimum"))