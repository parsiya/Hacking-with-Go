package main

import "fmt"

func main() {

	if var1 := 20; var1 > 10 {
		fmt.Println("Inside if:", var1)
	}
	// Cannot use the variable var1 here
}
