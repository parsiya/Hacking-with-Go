package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// generateStrings generated n strings and sends them to channel.
// Channel is closed when string generation is done.
func generateStrings(n int, c chan<- string) {

	// Close channel when done
	defer close(c)
	// Generate strings
	for i := 0; i < n; i++ {
		c <- fmt.Sprintf("String #%d", i)
	}
}

// consumeString reads strings from channel and prints them.
func consumeString(s string) {
	// Decrease waitgroup's counter by one
	defer wg.Done()
	fmt.Printf("Consumed %s\n", s)
}

func main() {
	// Create channel
	c := make(chan string)
	// Generate strings
	go generateStrings(10, c)

	for {
		select {
		// Read from channel
		case s, ok := <-c:
			// If channel is closed stop processing and return
			if !ok {
				// Wait for all goroutines to finish
				wg.Wait()
				// Return
				fmt.Println("Processing finished")
				return
			}
			// Increase wg counter by one for each goroutine
			// Note this is happening inside main before spawning the goroutine
			wg.Add(1)
			// Consume the string
			go consumeString(s)
		}
	}
}
