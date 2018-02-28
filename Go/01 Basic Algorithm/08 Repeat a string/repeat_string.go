package main

import "fmt"

func repeatStringNumTimes(str string, num int)  string {
	repeat := ""
	if num <= 0 {
		return repeat
	}
	for i := 0; i < num; i++ {
		repeat += str + "\n"
	}
	return repeat
}

func main() {
	fmt.Println(repeatStringNumTimes("abc", 3))
}