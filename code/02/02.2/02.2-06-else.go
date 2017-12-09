package main

import "fmt"

func main() {

	if var1 := 20; var1 > 100 {
		fmt.Println("Inside if:", var1)
	} else {
		// Can use var1 here
		fmt.Println("Inside else:", var1)
	}
	// Cannot use var1 here
}
