package main

import "fmt"

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
				fmt.Println("Processing finished")
				return
			}
			// Consume the string read from channel
			go consumeString(s)
		}
	}
}
