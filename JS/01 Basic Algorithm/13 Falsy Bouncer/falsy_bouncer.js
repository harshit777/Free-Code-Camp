function bouncer(arr) {
    return arr.filter(Boolean) ;
    //Check out the above code, the 'Boolean' that we pass if of the Object type and not the data type of true or false
    //check out std built in objects -> Boolean
    //arr.filter() creates a new array with the filtered data which we pass through
}

bouncer([7, "ate", "", false, 9, NaN, undefined, 0, null]);
