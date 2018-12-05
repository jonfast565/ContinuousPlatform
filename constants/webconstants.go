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
	MaxIdleConnections = 50
)

var DefaultTransport = &http.Transport{
	MaxIdleConnsPerHost: MaxIdleConnections,
}
