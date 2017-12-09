package main

import "fmt"

func main() {

	// Slice literal of type struct, the underlying array is created automatically
	sliceStruct := []struct {
		a, b int
	}{
		{1, 2},
		{3, 4},
		{5, 6}, // need this comma in the end otherwise it will not work
	}

	fmt.Println(sliceStruct)
}
