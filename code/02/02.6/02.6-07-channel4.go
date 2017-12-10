package main

import "fmt"

func main() {

	fourChan := make(chan int, 2)

	close(fourChan)

	i2 := 10
	fmt.Println("i2 before reading from closed channel", i2) // 10
	i2, ok := <-fourChan
	fmt.Printf("i2: %d - ok: %t", i2, ok) // i2: 10 - ok: false
}
