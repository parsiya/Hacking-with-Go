package main

import "fmt"

func main() {

	fourChan := make(chan int)

	go func() {
		// Send i1 to channel
		i1 := 10
		fourChan <- i1 // fourChan <- 10
		fmt.Printf("Sent %d to channel\n", i1)
	}()

	go func() {
		// Receive int from channel
		i2 := <-fourChan
		fmt.Println(i2)
		fmt.Printf("Received %d from channel\n", i2)
	}()

	// Wait for goroutines to finish
	fmt.Scanln()
	fmt.Println("Main Finished!")
}
