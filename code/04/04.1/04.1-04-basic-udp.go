package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"time"
)

var (
	host, port string
)

func init() {
	flag.StringVar(&port, "port", "80", "target port")
	flag.StringVar(&host, "host", "example.com", "target host")
}

func main() {

	// Converting host and port to host:port
	t := net.JoinHostPort(host, port)

	// Create a connection to server with 5 second timeout
	conn, err := net.DialTimeout("udp", t, 5*time.Second)
	if err != nil {
		panic(err)
	}

	// Write the GET request to connection
	// Note we are closing the HTTP connection with the Connection: close header
	// Fprintf writes to an io.writer
	req := "UDP PAYLOAD"
	fmt.Fprintf(conn, req)

	// Another way to do it to directly write bytes to conn with conn.Write
	// However we must first convert the string to bytes with []byte("string")
	// reqBytes := []byte(req)
	// conn.Write(reqBytes)

	// Reading the response

	// Create a new reader from connection
	connReader := bufio.NewReader(conn)

	// Read until a null byte (not safe in general)
	// Response will not be completely read if it has a null byte
	if status, err := connReader.ReadString(byte(0x00)); err != nil {
		fmt.Println(err)
		fmt.Println(status)
	}
}
