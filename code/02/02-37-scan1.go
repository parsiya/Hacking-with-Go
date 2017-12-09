package main

import "fmt"

func main() {

	var s string
	n, err := fmt.Scanln(&s)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Entered %d word(s): %s", n, s)
}
