package main

import (
	"fmt"
)


func sliceStrings(str string, start int , end int) string {
	return str[start:end]
}

func concatStrings(str1 string, str2 string) string {
	return str1 + " "+ str2
}

func main() {
	var myName string = "Fardin Hakimi"
	fmt.Printf("\nMy name is %v\t", myName)

	var sliced string = sliceStrings(myName, 0, 4)
	var joined string = concatStrings(myName, "Software Engineer")

	fmt.Printf("\n %v",sliced)
	fmt.Printf("\n %v\n",joined)

	if len(sliced) < 5 {
		fmt.Println("Sliced string is less than 5 chars in length")
	} else {
		fmt.Println("Sliced string has atleast 5 chars")
	}

	var multiLineStr string =
		`
			This is a multiline string
			line 1 
			line2 
			line 3
		`

	fmt.Println(multiLineStr)
}