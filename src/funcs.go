package main

import (
	"fmt"
)

func beforeMain() string {
	return "I am %d years old.\n"
}

func main() {
	fmt.Printf(beforeMain(), afterMain())
	firstName, lastName, age := returns2Strings()
	fmt.Printf("I am %v %v and I am %v years old.\n", firstName, lastName, age)
}

func afterMain() int {
	return 10
}

func returns2Strings() (string, string, float64) {
	return "Hayah", "Hakimi", 3.3
}


