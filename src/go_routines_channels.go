package main

import (
	"fmt"
	"time"
)

func main() {

	// Create a channel of ints
	ch := make(chan int)

	// Create a go routine to push a few numbers to the channel

	go func() {
		// Push numbers into the channel
		for i := 1; i <= 3; i++ {
			fmt.Printf("Pushing %d to the channel from routine 1 \n", i)
			ch <- i
			// Sleep for 1 second
			time.Sleep(time.Second)
		}
	}()

	go func() {
		// Push numbers into the channel
		for i := 4; i <= 6; i++ {
			fmt.Printf("Pushing %d to the channel from routine 2 \n", i)
			ch <- i
			// Sleep for 1 second
			time.Sleep(time.Second)
		}
		// Close the channel
		close(ch)
	}()

	// Read all numbers from the channel

	for i := 1; i <=3; i++ {
		val := <-ch
		fmt.Printf("Reading %d from the channel\n", val)
	}

	// Use the range to iterate over the channel

	for i := range ch {
		fmt.Printf("Recieved %d \n", i)
	}
}