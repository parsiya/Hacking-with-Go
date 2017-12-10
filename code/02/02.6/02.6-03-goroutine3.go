package main

import "fmt"

func main() {

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Printing from %d\n", 0)
		}
	}()

	// Wait for a keypress
	fmt.Scanln()
	fmt.Println("Main finished!")
}
