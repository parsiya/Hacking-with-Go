package main

import "fmt"

var emptyInterface interface{}

type Tuple struct {
	A, B int
}

func main() {

	// Use int
	int1 := 10
	emptyInterface = int1
	fmt.Println(emptyInterface) // 10

	// Use float
	float1 := 1.2345
	emptyInterface = float1
	fmt.Println(emptyInterface) // 1.2345

	// Use custom struct
	tuple1 := Tuple{5, 5}
	emptyInterface = tuple1
	fmt.Println(emptyInterface) // {5 5}
}
