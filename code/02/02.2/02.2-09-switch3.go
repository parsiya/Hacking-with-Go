package main

import (
	"fmt"
	"math/rand" // This is not cryptographically secure!
	"time"
)

func main() {
	// Seeding rand
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Choosing a random number:")

	num := rand.Intn(100)
	switch {
	case num < 50:
		fmt.Println("Less than 50")
	default:
		fmt.Println("More than 50")
	}
}
