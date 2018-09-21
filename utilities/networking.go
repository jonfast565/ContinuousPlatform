package utilities

import "strconv"

func getLocalPort(port int) string {
	return ":" + strconv.Itoa(port)
}
