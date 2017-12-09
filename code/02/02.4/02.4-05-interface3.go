package main

import "fmt"

func main() {
	var interface1 interface{} = 1234.5

	// Only print f1 if cast was successful
	if f1, ok := interface1.(float64); ok {
		fmt.Println("Float")
		fmt.Println(f1) // 1234.5
	}

	f2 := interface1.(float64)
	fmt.Println(f2) // 1234.5 No panic but not recommended

	// This will trigger a panic
	// i1 = interface1.(int)

	i2, ok := interface1.(int) // No panic
	fmt.Println(i2, ok)        // 0 false
}
