package networking

import (
	"fmt"
	"os"
	"strconv"
)

func GetLocalPort(port int) string {
	return ":" + strconv.Itoa(port)
}

func GetHostPortCombo(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}

func GetMyHostName() (string, error) {
	return os.Hostname()
}
