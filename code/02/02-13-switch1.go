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

	switch num := rand.Intn(3); num {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("3")
	}
}
