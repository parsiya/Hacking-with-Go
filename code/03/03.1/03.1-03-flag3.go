package main

import (
	"flag"
	"fmt"
)

// Declaring flag variables
var (
	ip      string
	port    int
	verbose bool
)

func init() {
	// Declaring flags
	// Remember, flag methods return pointers
	flag.StringVar(&ip, "ip", "127.0.0.1", "target IP")

	flag.IntVar(&port, "port", 8080, "Port")

	flag.BoolVar(&verbose, "verbose", false, "verbosity")
}

func main() {

	// Parsing flags
	flag.Parse()

	// Hack IP:port
	fmt.Printf("Hacking %s:%d!\n", ip, port)

	// Display progression if verbose flag is set
	if verbose {
		fmt.Printf("Pew pew!\n")
	}
}
