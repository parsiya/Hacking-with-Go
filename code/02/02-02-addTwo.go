package main

import "fmt"

func addTwo(x int, y int) (int, int) {
	return x + 2, y + 2
}

func main() {
	fmt.Println(addTwo(10, 20))
}
