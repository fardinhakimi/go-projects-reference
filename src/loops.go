package main 

import (
	"fmt"
)

func processSomeData() string {
	return " The number is %v\n"
}

func main() {

	for i :=0; i <= 10; i++ {

		if i == 5 {
			continue
		}

		if i == 10 {
			break
		}

		var str = processSomeData()
		fmt.Printf(str, i)
	}

	for i :=1; i <= 20; i++ {

		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("Fizz Buzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}