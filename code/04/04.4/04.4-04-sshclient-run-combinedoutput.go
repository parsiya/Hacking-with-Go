// SSH login with user/pass. Run a command with CombinedOutput and print output.

package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	// Importing crypto/ssh
	"golang.org/x/crypto/ssh"
)

var (
	username, password, serverIP, serverPort, command string
)

// Read flags
func init() {
	flag.StringVar(&serverPort, "port", "22", "SSH server port")
	flag.StringVar(&serverIP, "ip", "127.0.0.1", "SSH server IP")
	flag.StringVar(&username, "user", "", "username")
	flag.StringVar(&password, "pass", "", "password")
	flag.StringVar(&command, "cmd", "", "command to run")
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
		// This callback function validates the server.
		// Danger! We are ignoring host info
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
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

	// Close the session when main returns
	defer session.Close()

	// Run a command with CombinedOutput
	o, err := session.CombinedOutput(command)
	if err != nil {
		fmt.Println("Error running command", err)
	}

	fmt.Printf("Output:\n%s", o)
}
