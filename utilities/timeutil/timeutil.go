// Small time formatting package
package timeutil

import "time"

// Gets the RFC 850 time (the most printable format)
func GetCurrentTime() string {
	return time.Now().Format(time.RFC850)
}
