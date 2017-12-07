package main

import "fmt"

func main() {
	defer fmt.Println("This runs after main")

	fmt.Println("Main ended")
}
