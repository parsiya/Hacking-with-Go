package main

import (
	"log"
	"os"
)

func main() {

	a, b := 10, 20

	// New logger will output to stdout with prefix "Log1: " and flags
	// Note the space in prefix
	myLog := log.New(os.Stdout, "Log1: ", log.Ldate|log.Lshortfile)

	myLog.Print("Use Print to log.")
	myLog.Println("Ditto for Println.")
	myLog.Printf("Use Printf and format strings. %d + %d = %d", a, b, a+b)
}
