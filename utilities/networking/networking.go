// A simple networking package
// Used for getting host information
package networking

import (
	"fmt"
	"os"
	"strconv"
)

// Gets the local port in a format that can be used by the Gorilla MUX library
func GetLocalPort(port int) string {
	return ":" + strconv.Itoa(port)
}

// Gets the host port combo
func GetHostPortCombo(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}

// Wrapper: gets the hostname from the OS
func GetMyHostName() (string, error) {
	return os.Hostname()
}
