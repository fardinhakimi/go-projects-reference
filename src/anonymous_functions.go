package main

import (
	"fmt"
)

func main() {

	func(name string) {
		fmt.Println(name)
	}("My name is fardin!")
}
