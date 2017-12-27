package main

import (
	"flag"
	"fmt"
	"io"
	"net"
)

var (
	bindIP, bindPort, destIP, destPort string
)

func init() {
	flag.StringVar(&bindPort, "bindPort", "12345", "bind port")
	flag.StringVar(&bindIP, "bindIP", "127.0.0.1", "bind IP")
	flag.StringVar(&destPort, "destPort", "12345", "bind port")
	flag.StringVar(&destIP, "destIP", "127.0.0.1", "bind IP")
}

// readSocket reads data from socket if available and passes it to channel
func readSocket(conn net.Conn, c chan<- []byte) {

	// Create a buffer to hold data
	buf := make([]byte, 2048)
	// Store remote IP:port for logging
	rAddr := conn.RemoteAddr().String()

	for {
		// Read from connection
		n, err := conn.Read(buf)
		// If connection is closed from the other side
		if err == io.EOF {
			// Close the connction and return
			fmt.Println("Connection closed from", rAddr)
			return
		}
		// For other errors, print the error and return
		if err != nil {
			fmt.Println("Error reading from socket", err)
			return
		}
		// Print data read from socket
		// Note we are only printing and sending the first n bytes.
		// n is the number of bytes read from the connection
		fmt.Printf("Received from %v: %s\n", rAddr, buf[:n])
		// Send data to channel
		c <- buf[:n]
	}
}

// writeSocket reads data from channel and writes it to socket
func writeSocket(conn net.Conn, c <-chan []byte) {

	// Create a buffer to hold data
	buf := make([]byte, 2048)
	// Store remote IP:port for logging
	rAddr := conn.RemoteAddr().String()

	for {
		// Read from channel and copy to buffer
		buf = <-c
		// Write buffer
		n, err := conn.Write(buf)
		// If connection is closed from the other side
		if err == io.EOF {
			// Close the connction and return
			fmt.Println("Connection closed from", rAddr)
			return
		}
		// For other errors, print the error and return
		if err != nil {
			fmt.Println("Error writing to socket", err)
			return
		}
		// Log data sent
		fmt.Printf("Sent to %v: %s\n", rAddr, buf[:n])
	}
}

// forwardConnection creates a connection to the server and then passes packets
func forwardConnection(clientConn net.Conn) {

	// Converting host and port to destIP:destPort
	t := net.JoinHostPort(destIP, destPort)

	// Create a connection to server
	serverConn, err := net.Dial("tcp", t)
	if err != nil {
		fmt.Println(err)
		clientConn.Close()
		return
	}

	// Client to server channel
	c2s := make(chan []byte, 2048)
	// Server to client channel
	s2c := make(chan []byte, 2048)

	go readSocket(clientConn, c2s)
	go writeSocket(serverConn, c2s)
	go readSocket(serverConn, s2c)
	go writeSocket(clientConn, s2c)

}
func main() {

	flag.Parse()

	// Converting host and port to bindIP:bindPort
	t := net.JoinHostPort(bindIP, bindPort)

	// Listen for connections on BindIP:BindPort
	ln, err := net.Listen("tcp", t)
	if err != nil {
		// If we cannot bind, print the error and quit
		panic(err)
	}

	fmt.Printf("Started listening on %v\n", t)

	// Wait for connections
	for {
		// Accept a connection
		conn, err := ln.Accept()
		if err != nil {
			// If there was an error print it and go back to listening
			fmt.Println(err)

			continue
		}
		fmt.Printf("Received connection from %v\n", conn.RemoteAddr().String())

		go forwardConnection(conn)
	}
}
