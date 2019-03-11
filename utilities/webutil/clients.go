// Web related utilities
package webutil

import (
	"github.com/jonfast565/continuous-platform/constants"
	"net/http"
)

// Creates a new HttpClient
// Each new client uses the same transport as to effectively channel requests
// This is similar to the .NET use of HttpClient()
func NewHttpClient() *http.Client {
	client := &http.Client{
		Transport: constants.DefaultTransport,
		Timeout:   constants.ClientTimeout,
	}
	return client
}
