package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	sub1 *flag.FlagSet
	sub2 *flag.FlagSet

	sub1flag  *int
	sub2flag1 *string
	sub2flag2 int

	usage string
)

func init() {
	// Declare subcommand sub1
	sub1 = flag.NewFlagSet("sub1", flag.ExitOnError)
	// int flag for sub1
	sub1flag = sub1.Int("sub1flag", 0, "subcommand1 flag")

	// Declare subcommand sub2
	sub2 = flag.NewFlagSet("sub2", flag.ContinueOnError)
	// string flag for sub2
	sub2flag1 = sub2.String("sub2flag1", "", "subcommand2 flag1")
	// int flag for sub2
	sub2.IntVar(&sub2flag2, "sub2flag2", 0, "subcommand2 flag2")
	// Create usage
	usage = "sub1 -sub1flag (int)\nsub2 -sub2flag1 (string) -sub2flag2 (int)"
}

func main() {
	// If subcommand is not provided, print error, usage and return
	if len(os.Args) < 2 {
		fmt.Println("Not enough parameters")
		fmt.Println(usage)
		return
	}

	// Check the sub command
	switch os.Args[1] {

	// Parse sub1
	case "sub1":
		sub1.Parse(os.Args[2:])

	// Parse sub2
	case "sub2":
		sub2.Parse(os.Args[2:])

	// If subcommand is -h or --help
	case "-h":
		fallthrough
	case "--help":
		fmt.Printf(usage)
		return
	default:
		fmt.Printf("Invalid subcommand %v", os.Args[1])
		return
	}

	// If sub1 was provided and parse, print the flags
	if sub1.Parsed() {
		fmt.Printf("subcommand1 with flag %v\n", *sub1flag)
		return
	}

	// If sub2 was provided and parse, print the flags
	if sub2.Parsed() {
		fmt.Printf("subcommand2 with flags %v, %v\n", *sub2flag1, sub2flag2)
		return
	}
}
