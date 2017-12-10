package main

import "fmt"

func main() {

	c := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Printing from %d\n", 0)
		}

		// Send true to channel when we are done
		c <- true
	}()

	// Main will wait until it receives something from c
	<-c
}
