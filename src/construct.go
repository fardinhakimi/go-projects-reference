package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

type Square struct {
	Center Point
	Length int
}

func (sqr *Square) Move(dx int, dy int) {
	sqr.Center.X += dx
	sqr.Center.Y += dy
}

func (sqr *Square) Area() int {
	return sqr.Length * sqr.Length
}

func main() {

	point := Point { X: 1, Y: 1 }

	sqr := Square{ Center: point, Length: 2 }

	fmt.Printf(" X: %v , Y: %v \n", sqr.Center.X, sqr.Center.Y)

	sqr.Move(9,10)

	fmt.Printf(" X: %v , Y: %v \n", sqr.Center.X, sqr.Center.Y)

	fmt.Printf("Area: %v \n", sqr.Area())
}