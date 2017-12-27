package main

import (
	"flag"
	"fmt"
	"io"
	"net"
)

var (
	bindHost, bindPort string
)

func init() {
	flag.StringVar(&bindPort, "port", "12345", "bind port")
	flag.StringVar(&bindHost, "host", "127.0.0.1", "bind host")
}

// handleConnectionNoLog echoes everything back without logging (easiest)
func handleConnectionNoLog(conn net.Conn) {

	rAddr := conn.RemoteAddr().String()
	defer fmt.Printf("Closed connection from %v\n", rAddr)

	// This will accomplish the echo if we do not want to log
	io.Copy(conn, conn)
}

func main() {

	flag.Parse()

	// Converting host and port to host:port
	t := net.JoinHostPort(bindHost, bindPort)

	// Listen for connections on BindIP:BindPort
	ln, err := net.Listen("tcp", t)
	if err != nil {
		// If we cannot bind, print the error and quit
		panic(err)
	}

	// Wait for connections
	for {
		// Accept a connection
		conn, err := ln.Accept()
		if err != nil {
			// If there was an error, print it and go back to listening
			fmt.Println(err)
			continue
		}

		fmt.Printf("Received connection from %v\n", conn.RemoteAddr().String())

		// Spawn a new goroutine to handle the connection
		go handleConnectionNoLog(conn)
	}
}
