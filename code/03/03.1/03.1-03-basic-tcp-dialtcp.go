// Basic TCP client using TCPDial and TCP specific methods
package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

const (
	TargetHost = "example.com"
	TargetPort = 80
)

// CreateTCPAddr converts host and port to *TCPAddr
func CreateTCPAddr(target string, port int) (*net.TCPAddr, error) {
	return net.ResolveTCPAddr("tcp", target+":"+strconv.Itoa(port))
}

func main() {

	// Converting host and port
	a, err := CreateTCPAddr(TargetHost, TargetPort)
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
