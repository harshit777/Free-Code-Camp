function chunkArrayInGroups(arr, size) {

    var chunckedArray = [];
    
    while (arr.length) {
        chunckedArray.push(arr.splice(0, size))
        //Splice removes the elements from the array
    }
    return chunckedArray
}

chunkArrayInGroups(["a", "b", "c", "d"], 2);
