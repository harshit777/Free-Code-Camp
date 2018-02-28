package main

import "fmt"

func isPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var word string
	fmt.Printf("Enter a word: ")
	fmt.Scanln(&word)
	fmt.Println("Is", word, "a palindrome:", isPalindrome(word))
}
