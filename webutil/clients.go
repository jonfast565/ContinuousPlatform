package webutil

import (
	"../constants"
	"net/http"
)

func NewHttpClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: constants.MaxIdleConnections,
		},
		Timeout: constants.ClientTimeout,
	}
	return client
}
