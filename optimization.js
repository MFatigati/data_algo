const { string } = require("yargs");

let basketball_players = [
  {first_name: "Jill", last_name: "Huang", team: "Gators"},
  {first_name: "Janko", last_name: "Barton", team: "Sharks"},
  {first_name: "Wanda", last_name: "Vakulskas", team: "Sharks"},
  {first_name: "Jill", last_name: "Moloney", team: "Gators"},
  {first_name: "Luuk", last_name: "Watkins", team: "Gators"}
]
let football_players = [
  {first_name: "Hanzla", last_name: "Radosti", team: "32ers"},
  {first_name: "Tina", last_name: "Watkins", team: "Barleycorns"},
  {first_name: "Alex", last_name: "Patel", team: "32ers"},
  {first_name: "Jill", last_name: "Huang", team: "Barleycorns"},
  {first_name: "Wanda", last_name: "Vakulskas", team: "Barleycorns"}
]

function returnDoubles(arr1, arr2) {
  let seen = {};
  let both = [];
  arr1.forEach(player => {
    seen[player.first_name + " " + player.last_name] = true;
  })
  arr2.forEach(player => {
    if (seen[player.first_name + " " + player.last_name]) {
      both.push(player.first_name + " " + player.last_name);
    }
  })
  return both;
}

// console.log(returnDoubles(basketball_players, football_players));

function findMissing(arr) {
  let allNums = {};
  for (let i = 0; i < arr.length; i++) {
    allNums[i] = false;
  }
  arr.forEach(num => {
    allNums[num] = true;
  })
  for (let num in allNums) {
    if (allNums[num] === false) {
      return num;
    }
  }
}

// console.log(findMissing([2, 3, 0, 6, 1, 5]))
// console.log(findMissing([8, 2, 3, 9, 4, 7, 5, 0, 6]))

function findBestBuySell(arrStocks) {
  let currentBuy = arrStocks[0];
  let currentSale;
  let priorLowest = arrStocks[0];
  let profit = 0;
  for (let i = 0; i < arrStocks.length; i++) {
    if (arrStocks[i] < priorLowest) {
      priorLowest = arrStocks[i];
    }
    if (arrStocks[i] - priorLowest > profit) {
      currentBuy = priorLowest;
      currentSale = arrStocks[i];
      profit = arrStocks[i] - priorLowest;
    }
  }
  return [currentBuy, currentSale];
}

// console.log(findBestBuySell([10, 7, 5, 8, 11, 2, 6]));
// console.log(findBestBuySell([10, 7, 5, 8, 11, 2, 12]));
// console.log(findBestBuySell([10, 7, 1, 5, 11, 2, 12]));
// console.log(findBestBuySell([12, 7, 1, 5, 11, 2, 1]));

function highestProduct(arr) {
  let currentProd = 0;
  let lowestNeg = 0;
  let highestPos = 0;
  let toMultWith;
  arr.forEach(num => {
    if (Math.sign(num) === -1) {
      toMultWith = lowestNeg
      if (num < lowestNeg) {
        lowestNeg = num;
      }
    }
    if (Math.sign(num) === 1) {
      toMultWith = highestPos
      if (num > highestPos) {
        highestPos = num;
      }
    }
    if (num * toMultWith > currentProd) {
      currentProd = num * toMultWith;
    }
  })
  return currentProd;
}

//console.log(highestProduct([5, -10, -6, 9, 4]))

/**
 * 
 * 
 * Iterate over each number
 *  if value at i < priorlowest
 *    set prior lowest to value at i
 *  if value at i - priorLowest > profit
 *    set currentbuy to priorlowest
 *    set currentSale to value at i
 *  
 *
 * currentBuy 10
 * priorLowest 0
 * profit 0
 * 
 * currentBuy 10
 * currentSell 7
 * priorLowest 7
 * profit -3
 * 
 * currentBuy 7
 * currentSell 5
 * priorLowest 5
 * profit -2
 * 
 * currentBuy 5
 * currentSell 8
 * priorLowest 5
 * profit 3
 * 
 * 
 */

function sortTemps(arr) {
  let decimals = {
    // 0: 0,
    // 1: 0,
    // 2: 0,
    // 3: 0,
    // 4: 0,
    // 5: 0,
    // 6: 0,
    // 7: 0,
    // 8: 0,
    // 9: 0,
  }

  let outermost = {
    97: Object.assign({}, decimals),
    98: Object.assign({}, decimals),
    99: Object.assign({}, decimals)
  }
  
  arr.forEach(num => {
    let whole = Math.floor(num);
    let decimal = (num - Math.floor(num)).toFixed(1);
    outermost[whole][decimal] = outermost[whole][decimal] ? outermost[whole][decimal] + 1 : 1
  })

  let result = []

  //console.log(outermost[97]["0.1"])
  for (let outer in outermost) {

    for (let inner in outermost[outer]) {
      let timesToRepeat = outermost[outer][inner];
      while (timesToRepeat > 0) {
        result.push(Number(outer) + Number(inner));
        timesToRepeat--;
      }
    }
  }

  return result;
}

//console.log(sortTemps([98.6, 98.0, 97.1, 99.0, 98.9, 97.8, 98.5, 98.2, 98.0, 97.1]))

function longestConseqSeq(arr) {
  let seen = {};
  let longestSeq = 0;
  arr.forEach(num => {
    seen[num] = true;
    let currentlongestSeq = 1;
    let numToCheck = num;
    while (true) {
      if (seen[numToCheck - 1]) {
        currentlongestSeq++;
        numToCheck--;
      }
      if (!seen[numToCheck - 1]) break;
    }
    if (currentlongestSeq > longestSeq) {
      longestSeq = currentlongestSeq;
    }
    currentlongestSeq = 1;
    while (true) {
      if (seen[numToCheck + 1]) {
        currentlongestSeq++;
        numToCheck++;
      }
      if (!seen[numToCheck + 1]) break;
    }
    if (currentlongestSeq > longestSeq) {
      longestSeq = currentlongestSeq;
    }
  })
  return longestSeq;
}

console.log(longestConseqSeq([10, 5, 12, 3, 55, 30, 4, 11, 2]));
console.log(longestConseqSeq([19, 13, 15, 12, 18, 14, 17, 11]));