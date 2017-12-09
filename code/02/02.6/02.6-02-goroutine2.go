package main

import "fmt"

func Count(limit int) {
	for i := 0; i < limit; i++ {
		fmt.Printf("Counting: %d\n", i) // Now we know Printf
	}
}

func main() {

	go Count(100)
	go Count(100)

	// Wait for a keypress
	var s string
	fmt.Scanln(&s)
	fmt.Println("Main finished!")
}
