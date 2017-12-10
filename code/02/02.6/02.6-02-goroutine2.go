package main

import "fmt"

func PrintMe(t int, count int) {
	for i := 0; i < count; i++ {
		fmt.Printf("Printing from %d\n", t)
	}
}

func main() {

	go PrintMe(0, 10)

	// Wait for a keypress
	fmt.Scanln()
	fmt.Println("Main finished!")
}
