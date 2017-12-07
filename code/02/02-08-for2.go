package main

import "fmt"

func main() {
	// var sum int
	sum, i := 0
	for i < 20 { // while (i<20)
		sum += i
		i++
	}

	fmt.Println(sum)
}
