package main 


import (
	"fmt"
	"io"
)


type Capper construct {
	wtr io.Writer
}


func (c *Capper) Write(p []byte) (int, error) {
	return c.
}