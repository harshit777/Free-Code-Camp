package main

import (
	"fmt"
	"strings"
)

func mutation(slc, compare string) bool {
	//convert the case to lower for testing consistency
	target := strings.ToLower(slc)
	test := strings.ToLower(compare)
	//Loop through 'test'
	for _, v := range test {
		//Check if all the characters are present in 'target'
		if !strings.ContainsAny(target, string(v)) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(mutation("Alien", "lIn"))
}
