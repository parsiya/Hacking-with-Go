// SSH login with custom host verifier (only outputs host information).

package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"

	// Importing crypto/ssh
	"golang.org/x/crypto/ssh"
)

var (
	username, password, serverIP, serverPort string
)

// Read flags
func init() {
	flag.StringVar(&serverPort, "port", "22", "SSH server port")
	flag.StringVar(&serverIP, "ip", "127.0.0.1", "SSH server IP")
	flag.StringVar(&username, "user", "", "username")
	flag.StringVar(&password, "pass", "", "password")
}

// hostChecker returns a function to be used as callback for HostKeyCallback.
func hostChecker() ssh.HostKeyCallback {
	return printServerKey
}

// printServerKey prints server's info instead of checking it.
// It's of type HostKeyCallback
func printServerKey(hostname string, remote net.Addr, key ssh.PublicKey) error {
	// Just print everything
	fmt.Printf("Hostname: %v\nRemote address: %v\nServer key: %+v\n",
		hostname, remote, key)
	// Return nil so connection can continue without checking the server
	return nil
}

func main() {
	// Parse flags
	flag.Parse()

	// Check if username has been submitted - password can be empty
	if username == "" {
		fmt.Println("Must supply username")
		os.Exit(2)
	}

	// Create SSH config
	config := &ssh.ClientConfig{
		// Username
		User: username,
		// Each config must have one AuthMethod. In this case we use password
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		// Check the host
		HostKeyCallback: hostChecker(),
	}

	// Server address
	t := net.JoinHostPort(serverIP, serverPort)

	// Connect to the SSH server
	sshConn, err := ssh.Dial("tcp", t, config)
	if err != nil {
		fmt.Printf("Failed to connect to %v\n", t)
		fmt.Println(err)
		os.Exit(2)
	}

	// Create new SSH session
	session, err := sshConn.NewSession()
	if err != nil {
		fmt.Printf("Cannot create SSH session to %v\n", t)
		fmt.Println(err)
		os.Exit(2)
	}

	// Close the session
	defer session.Close()

	// For an interactive session we must redirect IO
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	input, err := session.StdinPipe()
	if err != nil {
		fmt.Println("Error redirecting session input", err)
		os.Exit(2)
	}

	// Setup terminal mode when requesting pty. You can see all terminal modes at
	// https://github.com/golang/crypto/blob/master/ssh/session.go#L56 or read
	// the RFC for explanation https://tools.ietf.org/html/rfc4254#section-8
	termModes := ssh.TerminalModes{
		ssh.ECHO: 0, // Disable echo
	}

	// Request pty
	// https://tools.ietf.org/html/rfc4254#section-6.2
	// First variable is term environment variable value which specifies terminal.
	// term doesn't really matter here, we will use "vt220".
	// Next are height and width: (40,80) characters and finall termModes.
	err = session.RequestPty("vt220", 40, 80, termModes)
	if err != nil {
		fmt.Println("RequestPty failed", err)
		os.Exit(2)
	}

	// Now we can start a remote shell
	err = session.Shell()
	if err != nil {
		fmt.Println("shell failed", err)
		os.Exit(2)
	}

	// Endless loop to capture commands
	for {
		io.Copy(input, os.Stdin)
	}
}
