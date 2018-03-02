function destroyer(arr) {
    //Split the array
    var args = Array.prototype.slice.call(arguments);

    //Loop through and delete same elements
    for (var i = 0; i < arr.length; i++) {
        for (var j = 0; j < args.length; j++) {
            if (arr[i] === args[j]) {
                delete arr[i];
            }
        }
    }

    //Simply filtering the array for null values using the Boolean object
    return arr.filter(Boolean);
}

destroyer([1, 2, 3, 1, 2, 3], 2, 3);
