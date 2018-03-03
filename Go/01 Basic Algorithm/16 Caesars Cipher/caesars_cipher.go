package main

import (
	"fmt"
	"strings"
)

func main() {
	//Thanks to the Go authors for writing this algorithm in their examples
	//Small task for Go author, big learning for newbies !
	rot13 := func(r rune) rune {
		//Note that in Go all strings are a read only slice of bytes
		//Slices can be easily converted into rune for carrying out operations
		//A rune is an alias of int32
		switch {
		//Now this is a nice syntax, note that we skip special characters & numbers for rot13
		case r >= 'A' && r <= 'Z':
			//Go authors are quite smart, imagine what happens if we have alphabet 's'
			//s + 13 will not lead us to the desired character, so we need to do s - 13

			//English has 26 alphabets, that means for the first 13 alphabets we add 13
			//and subtract 13 for the next half !

			//Go authors achieve this by taking the first alphabet 'A' or 'a' as reference
			//and then adding the required num to get our desired alphabet

			//(r-'A'+13) -> gives the number to add and modulus 26 gives it in reference to the first alphabet
			//e.g suppose rune of 'A' = 0, and we pass 'A' for rot13 as 'r' then =>
			//Step 1.    (r-'A'+13) = (0-0+13) = 13
			//Step 2.    13 % 26 = 13
			//Step 3.    'A' + 13 = 0 + 13 = 13
			//Step 4.     Suppose 13 is represented by the alphabet 'N', then the code returns 'N'

			//We save a couple of variables and if statements this way !
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	//We use the strings.Map a function from the Go standard library, chk out the doc at https://golang.org/pkg/strings
	fmt.Println(strings.Map(rot13, "SERR PBQR PNZC" ))
	fmt.Println(strings.Map(rot13, "SERR CVMMN!"))
	fmt.Println(strings.Map(rot13, "SERR YBIR?" ))
	fmt.Println(strings.Map(rot13, "GUR DHVPX OEBJA QBT WHZCRQ BIRE GUR YNML SBK." ))
}
