package main

import "fmt"

func slasher(slc []int, howMany int) (returnSlice []int) {
	return slc[howMany:]
}

func main() {
	slc := []int {1, 2, 3}
	fmt.Println(slasher(slc, 2))
}
