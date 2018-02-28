function mutation(arr) {
    var target = arr[0].toLowerCase();
    var compare = arr[1].toLowerCase();
    for (var i = 0; i < compare.length; i++) {
        if (target.indexOf(compare[i]) === -1) {
            return false;
        }
    }
    return true;
}

mutation(["hello", "hey"]);
