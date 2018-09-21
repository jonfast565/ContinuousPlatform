package utilities

import "time"

func getCurrentTime() string {
	return time.Now().Format(time.RFC850)
}