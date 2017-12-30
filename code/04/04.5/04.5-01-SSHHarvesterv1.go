// SSH Harvester version 1
// This is mainly a proof of concept during learning how to parse and verify
// SSH server certificates and host keys.
// Addresses should be in format of 'host:port'.
// Input file should have one address on each line and addresses provided to
// -targets/-t should be separated by commas
// -in and -targets are mutually exclusive, use one.
// -i, -in       string  input file
// -o, -out      string  output report file
// -t, -targets  string  addresses separated by comma
// -v, -verbose  string  print extra info
// See the blog post for discussion and techniques.

package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

const (
	// Usage constants
	mUsage = "SSH Harvester gathers and publishes info about SSH servers.\n" +
		"Addresses should be in format of 'host:port'.\n" +
		"Input file should have one address on each line " +
		"and addresses provided to -targets should be separated by commas.\n" +
		"-in and -targets are mutually exclusive, use one.\n"
	outUsage = "output report file"
	inUsage  = "input file"
	tUsage   = "addresses separated by comma"
	vUsage   = "print extra info"

	// Delimiter for host:port
	AddressDelim = ":"
	// // Delimiter for IPv6 addresses
	// IPv6Delim = "[]"

	// Log prefix - note the trailing space
	LogPrefix = "[*] "

	// Test SSH username/password - not really important
	TestUser     = "user"
	TestPassword = "password"

	// Timeout in seconds
	Timeout = 5 * time.Second
)

// Usage string
func usage() {
	usg := mUsage
	usg += fmt.Sprintf("\n  -i, -in\tstring\t%s", inUsage)
	usg += fmt.Sprintf("\n  -o, -out\tstring\t%s", outUsage)
	usg += fmt.Sprintf("\n  -t, -targets\tstring\t%s", tUsage)
	usg += fmt.Sprintf("\n  -v, -verbose\tstring\t%s", vUsage)
	usg += fmt.Sprintf("\n")

	fmt.Println(usg)
}

var (
	// Flags
	out     string  // output file
	in      string  // input file
	targets strList // addresses from command line
	verbose bool    // print extra information to stdout

	// Logger
	logSSH *log.Logger

	// Waitgroup for syncing discovery goroutines
	discoveryWG sync.WaitGroup
)

func init() {
	// Setup flags
	flag.StringVar(&out, "out", "", outUsage)
	flag.StringVar(&out, "o", "", outUsage)
	flag.StringVar(&in, "in", "", inUsage)
	flag.StringVar(&in, "i", "", inUsage)
	flag.Var(&targets, "targets", tUsage)
	flag.Var(&targets, "t", tUsage)
	flag.BoolVar(&verbose, "verbose", false, vUsage)
	flag.BoolVar(&verbose, "v", false, vUsage)

	// Set flag usage
	flag.Usage = usage

	// Parse flags
	flag.Parse()

	// Setting up logging
	logSSH = log.New(os.Stdout, LogPrefix, log.Ltime)

	// Check if we have enough arguments
	if len(os.Args) < 2 {
		flag.Usage()
		errorExit("not enough arguments", nil)
	}

	// Check if both in and targets are supported
	if (in != "") && (targets != nil) {
		errorExit("-in and -targets are mutually exclusive, use one", nil)
	}
}

// errorExit logs an error and then exits with status code 1.
func errorExit(m string, err error) {
	// If err is provided print it, otherwise don't
	if err != nil {
		logSSH.Fatalf("%v - stopping\n%v\n", m, err)
	}
	logSSH.Fatalf("%v - stopping\n", m)
}

// -----
// Custom flag type for -t (code re-used from flag section)
// Create a custom type from a string slice
type strList []string

// Implement String()
func (str *strList) String() string {
	return fmt.Sprintf("%v", *str)
}

// Implement Set(*strList)
func (str *strList) Set(s string) error {
	// If input was empty, return an error
	if s == "" {
		return errors.New("nil input")
	}
	// Split input by ","
	*str = strings.Split(s, ",")
	// Do not return an error
	return nil
}

// -----
// Struct to hold server data
type SSHServer struct {
	Address   string          // host:port
	Host      string          // IP address
	Port      int             // port
	IsSSH     bool            // true if server is running SSH on address:port
	Banner    string          // banner text, if any
	Cert      ssh.Certificate // server's certificate
	Hostname  string          // hostname
	PublicKey ssh.PublicKey   // server's public key
}

// NewSSHServer returns a new SSHServer with address, host and port populated.
// If address cannot be processed, an error will be returned.
func NewSSHServer(address string) (*SSHServer, error) {
	// Process address, return error if it's not in the correct format
	host, port, err := net.SplitHostPort(address)
	if err != nil {
		return nil, err
	}

	var s SSHServer

	s.Address = address
	s.Host = host
	s.Port, err = strconv.Atoi(port)
	if err != nil {
		return nil, err
	}
	// If port is not in (0,65535]
	if 0 > s.Port || s.Port > 65535 {
		return nil, errors.New(port + " invalid port")
	}
	return &s, nil
}

// discover connects to ip:port and attempts to make an SSH connection.
// If successful, some SSH properties will be populated (most importantly isSSH
// and isAlive).
func (s *SSHServer) discover() {
	// Release waitgroup after returning
	defer discoveryWG.Done()

	defer logSSH.Println("finished connecting to", s.Address)

	certCheck := &ssh.CertChecker{
		IsHostAuthority: hostAuthCallback(),
		IsRevoked:       certCallback(s),
		HostKeyFallback: hostCallback(s),
	}

	// Create SSH config
	config := &ssh.ClientConfig{
		// Test username and password
		User: TestUser,
		Auth: []ssh.AuthMethod{
			ssh.Password(TestPassword),
		},
		HostKeyCallback: certCheck.CheckHostKey,
		BannerCallback:  bannerCallback(s),
		Timeout:         Timeout, // timeout
	}

	logSSH.Println("starting SSH connection to ", s.Address)
	sshConn, err := ssh.Dial("tcp", s.Address, config)
	if err != nil {
		// If error contains "unable to authenticate", there's something there
		logSSH.Println("error ", err)
		return
	}

	// Close connection if we succeed (almost never happens)
	sshConn.Close()
}

type SSHServers []*SSHServer

// String converts []*SSHServer to JSON. If it cannot convert to JSON, it
// will convert each member to string using fmt.Sprintf("%+v").
func (servers *SSHServers) String() string {
	var report string
	// Try converting to JSON
	report, err := ToJSON(servers, true)
	// If cannot convert to JSON
	if err != nil {
		// Save all servers as string (this is not as good as JSON)
		for _, v := range *servers {
			report += fmt.Sprintf("%+v\n%s\n", v, strings.Repeat("-", 30))
		}
		return report
	}
	return report
}

// ToJSON converts input to JSON. If prettyPrint is set to True it will call
// MarshallIndent with 4 spaces.
// If your struct does not work here, make sure struct fields start with a
// capital letter. Otherwise they are not visible to the json package methods.
// We could also rewrite this as a method for ([]*SSHServer).
func ToJSON(s interface{}, prettyPrint bool) (string, error) {
	var js []byte
	var err error

	// Pretty print if specified
	if prettyPrint {
		js, err = json.MarshalIndent(s, "", "    ") // 4 spaces
	} else {
		js, err = json.Marshal(s)
	}

	// Check for marshalling errors
	if err != nil {
		return "", nil
	}

	return string(js), nil
}

// -----

// Define custom type for IsHostAuthority
type HostAuthorityCallBack func(ssh.PublicKey, string) bool

// hostAuthCallback is the callbackfunction for IsHostAuthority. Without
// it, ssh.CertChecker will not work.
func hostAuthCallback() HostAuthorityCallBack {
	// Return true because we just want to make this work
	return func(p ssh.PublicKey, addr string) bool {
		return true
	}
}

// Create IsRevoked function callback type
type IsRevokedCallback func(cert *ssh.Certificate) bool

// certCallback processes the SSH certificate. It is piggybacked on the
// IsRevoked callback function. It must return false (or nil) to keep the
// connection alive.
func certCallback(s *SSHServer) IsRevokedCallback {

	return func(cert *ssh.Certificate) bool {
		// Grab the certificate
		s.Cert = *cert
		s.IsSSH = true

		// Always return false
		return false
	}
}

// hostCallback is the callback function for HostKeyCallback in SSH config.
// It can access hostname, remote address and server's public key.
func hostCallback(s *SSHServer) ssh.HostKeyCallback {
	return func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		s.Hostname = hostname
		s.PublicKey = key
		// Return nil because we want the connection to move forward
		return nil
	}
}

// bannerCallback is the callback function for BannerCallback in SSH config.
// Grabs server banner and stores it in the SSHServer object.
func bannerCallback(s *SSHServer) ssh.BannerCallback {
	return func(message string) error {
		// Store the banner
		s.Banner = message
		// Return nil because we want the connection to move forward
		return nil
	}
}

// -----
// Misc functions

// readTargetFile opens a file and attempts to read targets from it.
// Each target should on its own line and in the correct format.
func readTargetFile(file string) ([]string, error) {

	var adds []string

	// Open the file and read it
	f, err := os.Open(in)
	if err != nil {
		return nil, err
	}
	// Close file
	defer f.Close()

	// Read line by line and add addresses
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		adds = append(adds, scanner.Text())
	}

	// Catch scanner errors
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return adds, nil
}

// writeReport stores results to file. Preferably uses ToJSON. If it cannot,
// prints them as string with .
func writeReport(file string, servers SSHServers) error {

	outfile, err := os.Create(file)
	if err != nil {
		return err
	}
	defer outfile.Close()

	// Try to serialize servers
	report := servers.String()
	// Write serialized date to file
	_, err = outfile.WriteString(report)
	if err != nil {
		return err
	}
	return nil
}

// -----
func main() {

	// Unprocessed addresses
	var adds []string
	var err error

	// If input file is provided - read targets
	if in != "" {
		logSSH.Printf("opening file %s", in)
		adds, err = readTargetFile(in)
		if err != nil {
			errorExit("error opening file", err)
		}
	}

	// If -t is provided
	adds = append(adds, targets...)

	// var servers SSHServers
	var servers SSHServers

	// Create a new SSHServer for each address. It might be a bit more efficient
	// to do a ping test first and only add live/accessible servers but ICMP
	// might be blocked and most importantly we will not save much time.
	for _, add := range adds {
		// Create temporary server and process it
		ts, err := NewSSHServer(add)
		if err != nil {
			logSSH.Printf("could not process %v", err)
			continue
		}
		// Add new server to servers
		servers = append(servers, ts)
	}

	// If no acceptable addresses are found
	if len(servers) == 0 {
		logSSH.Fatalf("no valid address was provided - terminating")
	}

	logSSH.Println("starting discovery")
	for _, v := range servers {
		// Before each goroutine add 1 to waitgroup
		discoveryWG.Add(1)
		go v.discover()
	}

	// Wait for all discovery goroutines to finish
	discoveryWG.Wait()

	logSSH.Println("finished discovery")

	// Write to file
	if out != "" {
		logSSH.Println("started writing report to file")
		err := writeReport(out, servers)
		if err != nil {
			logSSH.Println("error writing report to file - ", err)
		}
		logSSH.Println("finished writing report to file")
	} else {
		logSSH.Println("no output file specified, printing results")
		fmt.Println(ToJSON(servers, true))
	}

	logSSH.Println("finished")
}
