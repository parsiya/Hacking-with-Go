package main

import (
	"log"
)

func main() {

	a, b := 10, 20

	log.Print("Use Print to log.")
	log.Println("Ditto for Println.")
	log.Printf("Use Printf and format strings. %d + %d = %d", a, b, a+b)
}
