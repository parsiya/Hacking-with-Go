package main

import "fmt"

func main() {
	// var sum int
	sum := 0
	for i := 0; i < 20; i++ {
		sum += i
	}

	fmt.Println(sum)
}
