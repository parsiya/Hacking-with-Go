package main

import "fmt"

func main() {

	fourChan := make(chan int, 2)

	// Send 10 to channel
	fourChan <- 10
	fmt.Printf("Sent %d to channel\n", 10)

	// Receive int from channel
	// We can also receive directly
	fmt.Printf("Received %d from channel\n", <-fourChan)
}
