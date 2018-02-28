package main

import (
	"fmt"
	"strings"
)

func findLongestWord(str string)  {
	var longestWord string
	//Splitting the string into slice of words
	words := strings.Fields(str)
	//Initialising a tracker to check for the longest word
	tracker := 0
	//Looping the slice for the longest word
	for i := 0; i < len(words); i++ {
		if len(words[i]) > tracker {
			tracker = len(words[i])
			longestWord = words[i]
		}
	}
	//Printing out the longest word
	fmt.Printf("Longest Word: %s with Length: %v\n", longestWord, tracker)
}

func main() {
	line := "The quick brown fox jumped over the lazy dog"
	findLongestWord(line)

	newLine := "What if we try a super-long word such as otorhinolaryngology"
	findLongestWord(newLine)
}
