package main

import (
	"fmt"
)

func main() {

	x :=11

	if x > 12 {
		fmt.Println("x is greater than 12")
	} else if x == 11 {
		fmt.Println("x is 11!")
	} else {
		fmt.Println("x is less or equal to 12")
	}

	switch x {
		case 1:
			fmt.Println("one")
		case 11: 
			fmt.Println("11")
		case 10:
			fmt.Println("10")
		default:
			fmt.Println("I do not know!")
	}
}