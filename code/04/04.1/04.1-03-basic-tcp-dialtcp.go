// Basic TCP client using TCPDial and TCP specific methods
package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

var (
	host, port string
)

func init() {
	flag.StringVar(&port, "port", "80", "target port")
	flag.StringVar(&host, "host", "example.com", "target host")
}

// CreateTCPAddr converts host and port to *TCPAddr
func CreateTCPAddr(target, port string) (*net.TCPAddr, error) {
	return net.ResolveTCPAddr("tcp", net.JoinHostPort(host, port))
}

func main() {

	// Converting host and port to TCP address
	a, err := CreateTCPAddr(host, port)
	if err != nil {
		panic(err)
	}

	// Passing nil for local address
	tcpConn, err := net.DialTCP("tcp", nil, a)
	if err != nil {
		panic(err)
	}

	// Write the GET request to connection
	// Note we are closing the HTTP connection with the Connection: close header
	// Fprintf writes to an io.writer
	req := "GET / HTTP/1.1\r\nHost: example.com\r\nConnection: close\r\n\r\n"
	fmt.Fprintf(tcpConn, req)

	// Reading the response

	// Create a scanner
	scanner := bufio.NewScanner(bufio.NewReader(tcpConn))

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
