package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func toTitleCase(str string)  {
	//Split the string into slice of words
	//Convert all the words to lower case for consistency
	words := strings.Fields(strings.ToLower(str))
	//Loop through the slice
	for i := 0; i < len(words); i++ {
		//Getting each word from the slice
		word := words[i]
		//Capitalize the first character ->
		//1. Convert the word in []rune of string type
		//2. It is encoded in utf8 format
		//3. Capitalize the first character of the rune and append the rest of the slice
		r, n := utf8.DecodeRuneInString(word)
		words[i] = string(unicode.ToUpper(r)) + word[n:]
	}
	fmt.Println(strings.Join(words, " "))
}

func main() {
	str := "I'm a liTtle tea pot"
	toTitleCase(str)
}

