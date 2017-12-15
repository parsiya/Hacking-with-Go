package main

import (
	"flag"
	"fmt"
)

func main() {

	// Declare flags
	// Remember, flag methods return pointers
	ipPtr := flag.String("ip", "127.0.0.1", "target IP")

	var port int
	flag.IntVar(&port, "port", 8080, "Port")

	verbosePtr := flag.Bool("verbose", false, "verbosity")

	// Parse flags
	flag.Parse()

	// Hack IP:port
	fmt.Printf("Hacking %s:%d!\n", *ipPtr, port)

	// Display progress if verbose flag is set
	if *verbosePtr {
		fmt.Printf("Pew pew!\n")
	}
}
