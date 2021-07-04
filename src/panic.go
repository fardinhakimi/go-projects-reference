package main

import (
	"fmt"
	"os"
)

func throwPanic(fileUrl string) {

	file, err := os.Open(fileUrl)

	if err != nil {
		panic(err)
	}

	defer file.Close()
	fmt.Println(" file opened!")
}


func main() {

	throwPanic("no")
	vals := []int { 1,2,3}
	v := vals[10]
	fmt.Println(v)
}