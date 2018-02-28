package main

import (
	"fmt"
	"sort"
)

func largestOfFour(nums [][]int)  {
	var largestNumSlice []int
	var largestNum int
	for _, num := range nums {
		for i := 0; i < len(num)-1; i++ {
			sort.Ints(num)
			largestNum = num[len(num)-1]
		}
		largestNumSlice = append(largestNumSlice, largestNum,)
	}
	fmt.Printf("Sorted Slice: %v\n", nums)
	fmt.Printf("Largest Number is: %v\n", largestNum)
	fmt.Printf("Largest num slice is: %v\n", largestNumSlice)
}

func main() {
	nums := [][]int {}
	row1 := []int {4, 5, 1, 3}
	row2 := []int {13, 27, 18, 26}
	row3 := []int {32, 35, 37, 39}
	row4 := []int {1000, 1001, 857, 1}

	nums = append(nums, row1)
	nums = append(nums, row2)
	nums = append(nums, row3)
	nums = append(nums, row4)
	fmt.Println("Original Slice:", nums)

	largestOfFour(nums)
}
