package main

import "fmt"

// Create a new type for []string
type StringSlice []string

// Define the method for StringSlice
func (x StringSlice) PrintSlice() {
	for _, v := range x {
		fmt.Println(v)
	}
}

func main() {

	// Create an array of strings with 3 members
	characters := [3]string{"Ender", "Petra", "Mazer"}

	// Create a StringSlice
	var allMembers StringSlice = characters[0:3]

	// Now we can call the method on it
	allMembers.PrintSlice()

	// Ender
	// Petra
	// Mazer

	// allMembers.PrintSlice()
	// allMembers.PrintSlice undefined (type []string has no field or method PrintSlice)
}
