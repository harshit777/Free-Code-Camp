function confirmEnding(str, target) {
    if (str.substr(-target.length) === target) {
        return true;
    } else {
        return false ;
    }
}
//Using a ternary operator
//str.substr(-target.length) === target ? true : false;

confirmEnding('Bastian', 'n');