package main

import "fmt"

// Define new interface
type MyPrinter interface {
	MyPrint()
}

// Define a type for int
type MyInt int

// Define MyPrint() fro MyInt
func (i MyInt) MyPrint() {
	fmt.Println(i)
}

// Define a type for float64
type MyFloat float64

// Define MyPrint() for MyFloat
func (f MyFloat) MyPrint() {
	fmt.Println(f)
}

func main() {

	// Define interface
	var interface1 MyPrinter

	f1 := MyFloat(1.2345)
	// Assign a float to interface
	interface1 = f1
	// Call MyPrint() on interface
	interface1.MyPrint() // 1.2345

	i1 := MyInt(10)
	// Assign an int to interface
	interface1 = i1
	// Call MyPrint() on interface
	interface1.MyPrint() // 10
}
