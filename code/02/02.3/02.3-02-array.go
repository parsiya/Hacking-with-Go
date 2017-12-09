package main

import "fmt"

func main() {

	var a [5]int
	a[0] = 10
	a[4] = 20

	fmt.Println(a) // [10 0 0 0 20]

	// Array can be initialized during creation
	// characters[2] is empty
	characters := [3]string{"Ender", "Pentra"}

	fmt.Println(characters) // [Ender Pentra ]
}
