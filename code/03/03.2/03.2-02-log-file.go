package main

import (
	"log"
	"os"
)

func main() {

	// Create a file
	logFile, err := os.Create("log1.txt")
	if err != nil {
		panic("Could not open file")
	}

	// Close the file after main returns
	defer logFile.Close()

	a, b := 10, 20

	// We will not use the other options
	myLog := log.New(logFile, "", 0)

	myLog.Print("Use Print to log.")
	myLog.Println("Ditto for Println.")
	myLog.Printf("Use Printf and format strings. %d + %d = %d", a, b, a+b)
}
