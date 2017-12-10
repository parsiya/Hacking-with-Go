// This will not run

package main

import "fmt"

func main() {

	fourChan := make(chan int)

	i1 := 10

	// Send i1 to channel
	fourChan <- i1
	fmt.Printf("Sent %d to channel\n", i1)

	// Receive int from channel
	i2 := <-fourChan
	fmt.Printf("Received %d from channel\n", i2)
}
