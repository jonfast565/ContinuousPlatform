package timeutil

import "time"

func GetCurrentTime() string {
	return time.Now().Format(time.RFC850)
}

func GetCurrentSqlTime() string {
	return time.Now().Format("2006-01-02T15:04:05.999")
}
