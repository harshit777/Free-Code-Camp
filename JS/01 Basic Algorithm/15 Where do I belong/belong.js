
function getIndexToIns(arr, num) {
    arr.sort(function (a, b) { return a - b ;});
    console.log("Sorted array:" + arr);
    var index = 0;

    for (var i = 0; i < arr.length; i++) {
        if ((arr[i] < num) && (num <= arr[i+1])) {
            index = i+1;
        } else if (num > arr[arr.length-1]) {
            index = arr.length;
        }
    }
    return index;
}

getIndexToIns([40, 60], 50);