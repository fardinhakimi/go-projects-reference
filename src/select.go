package main

import (
	"fmt"
	"time"
)

func main() {

	ch1, ch2 := make(chan int), make(chan int)

	// Run a go routine to push a number into the first channel
	go func() {
		ch1 <- 10
	}()

	select {
	// In case we recieve a value from channel 1
	case val := <-ch1:
		fmt.Printf("Recieved ( %d ) from channel 1\n", val)
	// In case we recieve a value from channel 2
	case val := <-ch2:
		fmt.Printf("Recieved ( %d ) from channel 2\n", val)
	}

	//

	fmt.Println("------ IMPLEMENTINT TIME OUT WITH SELECT ------")

	out := make(chan float64)

	go func() {
		// sleep for 100 Millisecond
		time.Sleep(100 * time.Millisecond)
		// push value into the channel
		out <- 3.14
	}()

	select {
	case val := <-out:
		fmt.Printf("got %f \n", val)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("timeout")
	}
}
