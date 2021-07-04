package main 


import (
	"fmt"
	"math"
)



func findSquareRootOf(number float64) (float64, error) {

	if number < 0 {
		return 0.0, fmt.Errorf("The number is less than 0!")
	} else {
		return math.Sqrt(number), nil
	}
}

func main() {

	num, error := findSquareRootOf(-4)

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(num)
	}
}