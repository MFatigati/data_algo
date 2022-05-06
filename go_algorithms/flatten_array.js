function flatten(currentArr, depth = 1) {
    let temp = []
    
    if (depth === 0) {
        return currentArr
    }

    for (i = 0; i < currentArr.length; i++) {
        if (typeof currentArr[i] === "number") {
            temp.push(currentArr[i])
        } else {
            currentArr[i].forEach(elem => {
                temp.push(elem)
            })
        }
    }

    return flatten(temp, depth - 1)
}

console.log(flatten([1, 2, 3, [4, 5, [6, 7]], [8, 9]], 2))