package main 

import (
	"fmt"
)

func main() {

	names := []string { "hayah"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}