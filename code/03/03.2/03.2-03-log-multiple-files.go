package main

import (
	"bytes"
	"fmt"
	"io"
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

	// Create a second file
	logFile2, err := os.Create("log2.txt")
	if err != nil {
		panic("Could not open file2")
	}

	defer logFile2.Close()

	// Create a buffer
	var buflog bytes.Buffer

	multiW := io.MultiWriter(logFile, logFile2, &buflog, os.Stdout)

	a, b := 10, 20

	// Log to multiW
	myLog := log.New(multiW, "", 0)

	myLog.Print("Use Print to log.")
	myLog.Println("Ditto for Println.")
	myLog.Printf("Use Printf and format strings. %d + %d = %d", a, b, a+b)

	// Print buffer
	fmt.Println("Buffer:")
	fmt.Println(buflog.String())
}
