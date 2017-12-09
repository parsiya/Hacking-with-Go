package main

import (
	"fmt"
)

func main() {
	var a, b int = 20, 30
	// Need to convert a and b to float32 before the division
	var div float32 = float32(a) / float32(b)
	// Cast float32 to int
	var divInt = int(div)
	fmt.Println(div, divInt)
}
