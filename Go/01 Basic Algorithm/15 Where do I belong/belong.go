package main

import (
	"fmt"
	"sort"
)

func getIndexToIns(slc []int, num int)  {
	index := 0
	//Here we sort the slice
	sort.Ints(slc)
	fmt.Println("Sorted Slice:",slc)
	//Checking if slice is sorted
	if sort.IntsAreSorted(slc) {
		for i := 0; i < len(slc); i++ {
			 if num > slc[len(slc)-1] {
				index = len(slc)
			} else if slc[i] < num && num <= slc[i+1] {
				 index = i+1
			 }
		}
	}

	fmt.Println("Index position is:",index)
}

func main() {
	collection1 := []int {20,3,5}
	getIndexToIns(collection1, 19)

	collection2 := []int {10, 20, 30, 40, 50}
	getIndexToIns(collection2, 35)

	collection3 := []int {5, 3, 20, 3}
	getIndexToIns(collection3, 5)

	collection4 := []int {2, 5, 10}
	getIndexToIns(collection4, 15)
}
