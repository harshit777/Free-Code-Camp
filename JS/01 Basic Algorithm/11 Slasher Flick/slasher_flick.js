function slasher(arr, howMany) {
    // Splice returns the deleted elements
    arr.splice(0, howMany);
    //Returning the remaining array
    return arr;
}

slasher([1, 2, 3], 2);
