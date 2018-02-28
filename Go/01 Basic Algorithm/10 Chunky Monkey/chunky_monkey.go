//Now this algorithm had me breaking brains for an entire day !!!!
//A simple slicing error was breaking the code day long.
//Solutions on FCC forum & Google were not up to my taste of simplicity & some I couldn't understand
//so had to figure out my way, it may not be the fastest, but I prefer stability & simplicity over speed !
package main

import "fmt"

func chunkSliceInGroups(slc []rune, size int) (chunckedSlice [][]rune) {
	//Declare a temp []rune to hold our elements
	temp := []rune{}
	//Copying into a new variable as we don't want to play with the original slice
	tempSlc := slc
	//No of []slice to be added to the return [][]rune
	noOfSubSlice := len(slc) / size
	//As we know how many []slice are to be added, we iterate only till noOfSubSlices
	for j := 0; j < noOfSubSlice; j++ {
		//Load the temp slice till its length = 'size' of the original slice received
		temp = tempSlc[:size]
		//We now have a perfect temp slice which can be added to the return [][]rune
		chunckedSlice = append(chunckedSlice, temp)
		//Once added, we re-slice the original slice and delete the added elements
		tempSlc = tempSlc[size:]
		//Now flush the temp as it holds the above loaded elements
		temp = []rune{}
	}
	//This is where the last magic happens
	//After deleting the equal size slices we check there is any data left in the original slice
	//noOfSubSlice gives an integer, so we have equal size 'temp' slices of length = 'size'
	//So if the length of received slice is not perfectly divisible by 'size', some elements would
	//remain in the original slice, which then has to be pushed into the return [][]rune
	if len(tempSlc) != 0 {
		//Simpling loading temp with remaining elements
		temp = tempSlc[:]
		//Add the temp to the return [][]rune
		chunckedSlice = append(chunckedSlice, temp)
	}
	//Happily returning our perfect multi dimension slice
	return chunckedSlice

	//Didn't understand the code ? Try watching it after deleting the comments...
}

func main() {
	//Golang is a strictly typed language, so declaring the slice type as rune,
	//So that we can pass characters as well as ints
	arr1 := []rune{'a', 'b', 'c', 'd'}
	arr2 := []rune{0, 1, 2, 3, 4, 5}
	arr3 := []rune{0, 1, 2, 3, 4, 5}
	arr4 := []rune{0, 1, 2, 3, 4, 5}
	arr5 := []rune{0, 1, 2, 3, 4, 5, 6}
	arr6 := []rune{0, 1, 2, 3, 4, 5, 6, 7, 8}
	arr7 := []rune{0, 1, 2, 3, 4, 5, 6, 7, 8}

	//Printing the chuncked slices

	//Since the first slice is of character type, use printf otherwise the
	//int32 value equivalent to the rune will the printed.
	fmt.Printf("%c\n", chunkSliceInGroups(arr1, 2) )
	fmt.Println(chunkSliceInGroups(arr2, 3))
	fmt.Println(chunkSliceInGroups(arr3, 2))
	fmt.Println(chunkSliceInGroups(arr4, 4))
	fmt.Println(chunkSliceInGroups(arr5, 3))
	fmt.Println(chunkSliceInGroups(arr6, 4))
	fmt.Println(chunkSliceInGroups(arr7, 2))

}
