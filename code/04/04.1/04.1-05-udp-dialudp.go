package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strconv"
	"time"
)

var (
	host string
	port int
)

func init() {
	flag.IntVar(&port, "port", 80, "target port")
	flag.StringVar(&host, "host", "example.com", "target host")
}

// CreateUDPAddr converts host and port to *UDPAddr
func CreateUDPAddr(target string, port int) (*net.UDPAddr, error) {
	return net.ResolveUDPAddr("udp", target+":"+strconv.Itoa(port))
}

func main() {

	// Converting host and port
	a, err := CreateUDPAddr(host, port)
	if err != nil {
		panic(err)
	}

	// Create a connection with DialUDP
	connUDP, err := net.DialUDP("udp", nil, a)
	if err != nil {
		panic(err)
	}

	// Write the GET request to connection
	// Note we are closing the HTTP connection with the Connection: close header
	// Fprintf writes to an io.writer
	req := "UDP PAYLOAD"
	fmt.Fprintf(connUDP, req)

	// Reading the response

	// Create a scanner
	scanner := bufio.NewScanner(bufio.NewReader(connUDP))

	// Read from the scanner and print
	// Scanner reads until an I/O error
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	// Check if scanner has quit with an error
	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error", err)
	}
}
