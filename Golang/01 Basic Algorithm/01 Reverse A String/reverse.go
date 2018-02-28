//Reverse a string
package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
)

func Reverse(s string) string {
	rune := []rune(s)
	//Reversing the string
	for i, j := 0, len(rune)-1;i < j ; i, j = i + 1, j - 1  {
		rune[i], rune[j] = rune[j], rune[i]
	}
	return string(rune)
}

func main() {
	//Creating a buffered i/o reader
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the string you want to reverse: ")
	//Scanning the input from os.Stdin
	toReverse, err := reader.ReadString('\n')

	//Todo - Add condition to check if the input is a 'string'

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	reversedString := Reverse(toReverse)

	//Todo - Add a function isReverse() to check if it is reversed

	fmt.Println("The reversed string is:", reversedString)
}
