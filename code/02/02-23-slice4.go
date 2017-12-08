package main

import "fmt"

func main() {

	ints := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(ints)

	slice1 := ints[2:6]

	// len=4 and cap=4 (from 3rd item of the array until the end)
	printSlice(slice1)

	slice1 = ints[2:4]

	// len=2 but cap will remain 4
	printSlice(slice1)
}

// Copied from the tour
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
