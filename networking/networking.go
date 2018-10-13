package networking

import "strconv"

func GetLocalPort(port int) string {
	return ":" + strconv.Itoa(port)
}
