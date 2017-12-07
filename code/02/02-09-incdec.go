package main

import "fmt"

func main() {
	// var sum int
	sum, i := 0
	// This will not work
	sum = i++
	fmt.Println(sum)
}
