package main

import "fmt"

// Define a struct
type Tuple struct {
	A, B int
}

// Create a Stringer for Tuple
func (t Tuple) String() string {
	// Sprintf is similar to the equivalent in C
	return fmt.Sprintf("A: %d, B: %d", t.A, t.B)
}

func main() {

	tuple1 := Tuple{10, 10}
	tuple2 := Tuple{20, 20}
	fmt.Println(tuple1) // A: 10, B: 10
	fmt.Println(tuple2) // A: 20, B: 20
}
