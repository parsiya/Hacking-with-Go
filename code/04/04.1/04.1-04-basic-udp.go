// Replace TargetHost and TargetPort with a valid UDP server:port before running
package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"time"
)

const (
	TargetHost = "127.0.0.1"
	TargetPort = 80
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
