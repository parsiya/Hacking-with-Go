package main

import "fmt"

func addTwo2(x int, y int) (xPlusTwo int, yPlusTwo int) {
	xPlusTwo = x + 2
	yPlusTwo = y + 2

	return xPlusTwo, yPlusTwo
}

func main() {
	fmt.Println(addTwo2(20, 30))
}
