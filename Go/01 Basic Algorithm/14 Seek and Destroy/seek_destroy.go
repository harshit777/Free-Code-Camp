//Oh, holy devil, clearly not the best day and not the best code, but it works !
package main

import (
	"fmt"
)

//Not really the best code, but somehow wanted to make the code work

//Filters the slice with one filter at a time
func filterSliceElement(target []interface{}, seeker interface{}) []interface{} {
	makeFilter := target
	for i := 0; i < len(makeFilter); i++ {
		if makeFilter[i] == seeker {
			makeFilter = append(target[:i], makeFilter[i+1:]...)
		}
	}
	return makeFilter
}

//Filters the entire slice
func filterSlice(target, seeker []interface{}) []interface{}{
	// Note that both target & seeker are slice
	for _, v := range seeker {
		//Filter entire slice with first element
		target = filterSliceElement(target, v)
	}

	//Checking if filtered
	if isFiltered(target, seeker) {
		fmt.Println()
	} else {
		//Not the right way to do it, but just for the sake of getting the right output re-looping it
		for _, v := range seeker {
			target = filterSliceElement(target, v)
		}
	}
	return target
}

//Checking the slice for any filters
func isFiltered(target, seeker []interface{}) bool {
	isFilter := true
	for _, v := range seeker {
		for i := 0; i < len(target); i++ {
			if target[i] == v {
				isFilter = false
			}
		}
	}
	return isFilter
}

func main() {
	// Interfaces are really special to Go, notice that we can add multiple
	// data types with slice of type interface
	// Otherwise you won't be able to mix data types as Golang is strictly typed language
	data := []interface{}{"Orange", "Walnut", "Watermelon", "Monkey", "Mango", "Orange", 34, 34, 123, 123, "Watermelon", "Orange"}
	filterWith := []interface{}{"Orange", 34, 123, "Watermelon", "Ether"}

	fmt.Println("Filtered Slice:",filterSlice(data, filterWith))
}

//Not clearly the most efficient code, it just works, please let me know for any better solutions :)

//Tried posting it on stack overflow, and thanks to all the dick heads who asked me to post
//a 'real' question & down voting me, I had to come up with this solution, but hey, at least it works !

//To my dear stack overflow friends, any problem is a real problem if you can't solve it !