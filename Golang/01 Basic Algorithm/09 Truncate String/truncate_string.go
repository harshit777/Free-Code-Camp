package main

import "fmt"

func truncateString(str string, num int)  {
	var newStr string
	if  num >= len(str) {
		fmt.Println(str)
		return
	}
	if num <= 3 {
		newStr = str[:num] + "..."
		fmt.Println(newStr)
		//Alternative to skip var newStr
		//fmt.Println(str[:num] + "...")
		return
	}
	newStr = str[:num-3] + "..."
	fmt.Println(newStr)
}

func main() {
	truncateString("A-tisket a-tasket A green and yellow basket", 11)
	truncateString("Absolutely Longer", 2)
	truncateString("A-", 1)
}
