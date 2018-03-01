//TODO - Code Has A Bug!
package main

import (
	"fmt"
)

//Not really the best code, but somehow wanted to make the code work
func filterSliceElement(target []interface{}, seeker interface{}) []interface{} {
	markeFilter := target
	fmt.Println("Received Data:", target)
	fmt.Println("Element to delete:", seeker)
	for i := 0; i < len(markeFilter); i++ {
		if markeFilter[i] == seeker {
			markeFilter = append(target[:i], markeFilter[i+1:]...)
		}
	}
	fmt.Println("Returning this slice:", markeFilter)
	fmt.Println()
	return markeFilter
}

func filterSlice(target, seeker []interface{}) {
	// Note that both target & seeker are slice
	fmt.Println("Original Data:", target)
	fmt.Println("Element to filter:", seeker)

	for _, v := range seeker {
		fmt.Println("v is:", v)
		target = filterSliceElement(target, v)
	}
	fmt.Println("After filtering the elements:", target)
	fmt.Println()
}

func main() {
	// Interfaces are really special to Go, notice that we can add multiple
	// data types with slice of type interface
	// Otherwise you won't be able to mix data types as Golang is strictly typed language
	data1 := []interface{}{"Orange", "Walnut", "Watermelon", "Monkey", "Mango", "Orange", 34, 34, 123, 123, "Watermelon", "Orange"}
	filterWith := []interface{}{"Orange", 34, 123, "Watermelon", "Ether"}

	fmt.Println()
	fmt.Println("Filtered Data with single elemenet:", filterSliceElement(data1, 34))
	filterSlice(data1, filterWith)
}

//Expected Output - [Walnut Monkey Mango Orange]
// Received Output - [Walnut Monkey Mango 123 Orange]