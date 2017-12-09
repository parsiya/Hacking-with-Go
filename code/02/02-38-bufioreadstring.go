package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	// ReadString will read until first new line
	input, err := reader.ReadString('\n') // Need to pass '\n' as char (byte)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Entered %s", input)
}
