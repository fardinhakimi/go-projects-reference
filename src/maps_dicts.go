package main 

import (
	"fmt"
)

func main() {

	stocks := map[string]float64 { "amazon": 17.7, "facebook": 22.2 }

	var amazonStockValue float64 = stocks["amazon"]

	fmt.Printf("This amazon's current stock value is %v \n", amazonStockValue)

	value, exists := stocks["google"]

	if exists {
		fmt.Printf("The google's current stock price is %v", value)
	} else {
		fmt.Println("Google stock price is missing!")
	}

	stocks["google"] = 20.2

	fmt.Println(stocks)

	delete(stocks, "amazon")

	fmt.Println(stocks)
}