package main

import (
	"flag"
	"fmt"
)

func main() {
	// Set flag
	_ = flag.Int("flag1", 0, "flag1 description")
	// Parse all flags
	flag.Parse()
	// Enumererate flag.Args()
	for _, v := range flag.Args() {
		fmt.Println(v)
	}
	// Enumerate using flag.Arg(i)
	for i := 0; i < flag.NArg(); i++ {
		fmt.Println(flag.Arg(i))
	}
}
