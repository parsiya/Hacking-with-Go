package main

import "fmt"

// Directed write-only channel
func Sender(c chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("Sent", i)
		c <- i
	}
}

func Receiver(c <-chan int) {
	for i := range c {
		fmt.Println("Received", i)
	}
}

func main() {

	fourChan := make(chan int)

	go Sender(fourChan)
	go Receiver(fourChan)

	// Wait for goroutines to finish
	fmt.Scanln()
	fmt.Println("Main Finished!")
}
