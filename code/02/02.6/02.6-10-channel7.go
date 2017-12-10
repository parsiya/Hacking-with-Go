package main

import "fmt"

func main() {

	c := make(chan int, 2)

	for i := 0; i < 10; i++ {

		select {
		case c <- i:
			// If we can write to channel, send something to it
			fmt.Println("Sent to channel", i)
		case i2 := <-c:
			// If we can read from channel, read from it and print
			fmt.Println("Received from channel", i2)
		default:
			// This is run when nothing else can be done
		}
	}
}
