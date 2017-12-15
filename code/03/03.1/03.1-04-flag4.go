package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"sync"
)

// 1. Create a custom type from a string slice
type strList []string

// 2.1 implement String()
func (str *strList) String() string {
	return fmt.Sprintf("%v", *str)
}

// 2.2 implement Set(*strList)
func (str *strList) Set(s string) error {
	// If input was empty, return an error
	if s == "" {
		return errors.New("nil input")
	}
	// Split input by ","
	*str = strings.Split(s, ",")
	// Do not return an error
	return nil
}

// Declare flag variables
var (
	ip      strList
	port    strList
	verbose bool
)

var wg sync.WaitGroup

func init() {
	// Declare flags
	// Remember, flag methods return pointers
	flag.Var(&ip, "ip", "target IP")

	flag.Var(&port, "port", "Port")

	flag.BoolVar(&verbose, "verbose", false, "verbosity")
}

// permutations creates all permutations of ip:port and sends them to a channel.
// This is preferable to returing a []string because we can spawn it in a
// goroutine and process items in the channel while it's running. Also save
// memory by not creating a large []string that contains all permutations.
func permutations(ips strList, ports strList, c chan<- string) {

	// Close channel when done
	defer close(c)
	for _, i := range ips {
		for _, p := range ports {
			c <- fmt.Sprintf("%s:%s", i, p)
		}
	}
}

// hack spawns a goroutine that "hacks" each target.
// Each goroutine prints a status and display progres if verbose is true
func hack(target string, verbose bool) {

	// Reduce waitgroups counter by one when hack finishes
	defer wg.Done()
	// Hack the planet!
	fmt.Printf("Hacking %s!\n", target)

	// Display progress if verbose flag is set
	if verbose {
		fmt.Printf("Pew pew!\n")
	}
}

func main() {

	// Parse flags
	flag.Parse()

	// Create channel for writing and reading IP:ports
	c := make(chan string)

	// Perform the permutation in a goroutine and send the results to a channel
	// This way we can start "hacking" during permutation generation and
	// not create a huge list of strings in memory
	go permutations(ip, port, c)

	for {
		select {
		// Read a string from channel
		case t, ok := <-c:
			// If channel is closed
			if !ok {
				// Wait until all goroutines are done
				wg.Wait()
				// Print hacking is finished and return
				fmt.Println("Hacking finished!")
				return
			}
			// Otherwise increase wg's counter by one
			wg.Add(1)
			// Spawn a goroutine to hack IP:port read from channel
			go hack(t, verbose)
		}
	}
}
