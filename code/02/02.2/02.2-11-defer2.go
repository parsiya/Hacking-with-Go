package main

import "fmt"

func main() {
	num := 1
	defer fmt.Println("After main returns", num)

	num++
	fmt.Println("Inside main", num)
}
