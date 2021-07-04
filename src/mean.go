package main 

import (
	"fmt"
)

func calculateMean(number1 float64, number2 float64) float64 {
	var mean float64 = (number1 + number2) / 2.0
	return mean
}

func main() {
	var number2 float64 = 10
	var number1 float64
	number1 = 13
	var mean = calculateMean(number1, number2)
	fmt.Printf(" result: %v ", mean)
}