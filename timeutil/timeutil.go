package timeutil

import "time"

func GetCurrentTime() string {
	return time.Now().Format(time.RFC850)
}
