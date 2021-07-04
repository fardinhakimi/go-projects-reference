package main

import (
	"fmt"
	"hakimi.com/greetings"
)

func main() {
	message := greetings.Hello("hi there!")
	fmt.Println(message)
}
