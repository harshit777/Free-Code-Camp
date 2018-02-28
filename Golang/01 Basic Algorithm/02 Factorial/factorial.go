package main

import (
	"fmt"
)

func factorialize(n int) int {
	//Creating a limiting condition otherwise it will go into infinite recursion
	if n == 0 {
		return 1;
	}
	//Using a recursive function to get the factorial
	return n * factorialize(n - 1)
}

func main() {
	var toFactorial int
	fmt.Printf("Enter the number to find its factorial: ")
	fmt.Scan(&toFactorial)
	//Todo - Use a buffer i/o instead of Scan()

	fmt.Printf("The factorial of the number %v is: %v\n", toFactorial, factorialize(toFactorial))
}
