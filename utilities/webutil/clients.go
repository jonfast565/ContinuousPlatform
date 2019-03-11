package webutil

import (
	"github.com/jonfast565/continuous-platform/constants"
	"net/http"
)

func NewHttpClient() *http.Client {
	client := &http.Client{
		Transport: constants.DefaultTransport,
		Timeout:   constants.ClientTimeout,
	}
	return client
}
