package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

const (
	TargetHost = "127.0.0.1"
	TargetPort = 12345
)

// CreateAddress converts host and port to host:port.
func CreateAddress(target string, port int) string {
	// Running string(port) will return the char for 80 decimal or P
	// strconv.Itoa converts an int to string (e.g. 80 to "80")
	return target + ":" + strconv.Itoa(port)
}

func main() {

	// Converting host and port
	t := CreateAddress(TargetHost, TargetPort)

	// Create a connection to server
	conn, err := net.Dial("tcp", t)
	if err != nil {
		panic(err)
	}

	// Write the GET request to connection
	// Note we are closing the HTTP connection with the Connection: close header
	// Fprintf writes to an io.writer
	req := "GET / HTTP/1.1\r\nHost: example.com\r\nConnection: close\r\n\r\n"
	fmt.Fprintf(conn, req)

	// Another way to do it to directly write bytes to conn with conn.Write
	// However we must first convert the string to bytes with []byte("string")
	// reqBytes := []byte(req)
	// conn.Write(reqBytes)

	// Reading the response

	// Create a new reader from connection
	connReader := bufio.NewReader(conn)

	// Create a scanner
	scanner := bufio.NewScanner(connReader)

	// Combined into one line
	// scanner := bufio.NewScanner(bufio.NewReader(conn))

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
