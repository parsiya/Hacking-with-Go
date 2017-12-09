package main

import "fmt"

func printType(i interface{}) {
	// Do a type switch on interface
	switch val := i.(type) {
	// If an int is passed
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case float64:
		fmt.Println("float64")
	default:
		fmt.Println("Other:", val)
	}
}

func main() {
	printType(10)      // int
	printType("Hello") // string
	printType(156.32)  // float64
	printType(nil)     // Other: <nil>
	printType(false)   // Other: false
}
