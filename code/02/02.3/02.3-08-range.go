package main

import "fmt"

func main() {
	characters := [3]string{"Ender", "Petra", "Mazer"}
	for i, v := range characters {
		fmt.Println(i, v)
	}

	// 0 Ender
	// 1 Petra
	// 2 Mazer

	fmt.Println("-----------")

	// Only using index
	for i := range characters {
		fmt.Println(i, characters[i])
	}

	fmt.Println("-----------")

	// Ignoring index
	for _, v := range characters {
		// No non-elaborate way to get index here
		fmt.Println(v)
	}

	// Ender
	// Petra
	// Mazer
}
