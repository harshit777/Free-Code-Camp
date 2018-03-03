function rot13(str) {
    return str.split('')
        .map.call(str, function (char) {
        //takes the first character in the str array and returns its ASCII value
        x = char.charCodeAt(0);
        //Check if it lies between A to Z
        if (x < 65 || x > 90) {
            //Return the character as it is
            return String.fromCharCode(x);
        } else if (x < 78) {
            //If x lies till N then
            return String.fromCharCode(x + 13);
        }
        //Or reduce the x if character is more than N
        return String.fromCharCode(x - 13);
    }).join('');
}

rot13("SERR PBQR PNZC");

//After writing such code in JS it always make me prefer Golang, really Golang offers so much simplicity !
//Chaining objects or methods is a pain in JS
//Golang offers much better documentation and passing data in functions is nice and clean