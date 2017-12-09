package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	// Read bytes until the new line
	input, err := reader.ReadBytes('\n') // Need to pass '\n' as char (byte)
	if err != nil {
		panic(err)
	}

	// Print type of "input" and its value
	fmt.Printf("Entered type %T, %v\n", input, input)
	// Print bytes as string
	fmt.Printf("Print bytes as string with %%s %s", input)
}
