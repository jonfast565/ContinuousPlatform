package webutil

import (
	"../constants"
	"net/http"
)

func NewHttpClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 20,
		},
		Timeout: constants.ClientTimeout,
	}
	return client
}
