package constants

import (
	"net/http"
	"time"
)

var (
	PostMethod         = "POST"
	GetMethod          = "GET"
	DefaultScheme      = "http"
	ClientTimeout      = 60 * time.Minute
	MaxIdleConnections = 999999
)

var DefaultTransport = &http.Transport{
	MaxIdleConnsPerHost: MaxIdleConnections,
}
