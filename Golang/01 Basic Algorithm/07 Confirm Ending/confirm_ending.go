package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

//TODO - Refactor code according to methods
func confirmEnding(str1, str2 string) {
	//For consistency in testing, we convert the strings to lowercase
	//In golang uppercase & lowercase represent a different byte or UTF8 encoding
	//So we need to convert them in one format for testing
	str1Last := lastIndex(strings.ToLower(str1))
	str2Last := lastIndex(strings.ToLower(str2))
	if strings.Compare(str1Last, str2Last) == 0 {
		fmt.Println("Ends Match !")
	} else {
		fmt.Println("No Match!")
	}

	foo1 := strings.ToLower(string(str1[len(str1)-1]))
	foo2 := strings.ToLower(string(str2[len(str2)-1]))
	if strings.Compare( foo1, foo2 ) == 0 {
		fmt.Println("Foo Match !")
	} else {
		fmt.Println("Foo No Match !")
	}
}

func lastIndex(str string) string  {
	//In Golang strings are basically slice of bytes which can be indexed
	//But we are converting into rune in order to extract the last element
	fmt.Printf("2nd Char: %c\n", str[1])
	fmt.Printf("Last Character: %c\n", str[len(str)-1])

	fooLast := string(str[len(str)-1])
	fmt.Println("Foo is:", fooLast)

	r, _ := utf8.DecodeRuneInString(str)
	decodedRune := "A" + str[2:]
	fmt.Println("Decoded Rune:", decodedRune)

	fmt.Println("Decoded Rune In String:", string(r))
	fmt.Printf("Decoded Rune In Character: %c\n", r)

	rune := []rune(str)

	//Rune are an alias of type int32 encoded in UTF-8, so we need to convert them into string before comparing
	//As we want to compare as string, we put a cast and convert from rune to string and return the last character
	return string(rune[len(rune)-1])
}

func main() {
	confirmEnding("tejA", "lina")
}
