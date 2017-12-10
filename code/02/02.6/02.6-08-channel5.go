package main

import "fmt"

func main() {

	fourChan := make(chan int, 10)

	go func() {
		// Send 0-9 to channel
		for i := 0; i < 10; i++ {
			fourChan <- i
		}
	}()

	go func() {
		// Receive from channel
		for v := range fourChan {
			fmt.Println(v)
		}
	}()

	// Wait for goroutines to finish
	fmt.Scanln()
	fmt.Println("Main Finished!")
}
