package utilities

import (
	"io/ioutil"
	"net/http"
)

const (
	AuthorizationHeader              string = "Authentication"
	AuthorizationTypeString          string = "Basic"
	ContentTypeHeader                string = "Content-Type"
	ApplicationJsonHeaderContentType string = "application/json"
	OctetStreamHeaderContentType     string = "application/octet-stream"
)

func AddJsonHeader(request *http.Request) {
	request.Header.Add(ContentTypeHeader, ApplicationJsonHeaderContentType)
}

func AddOctetHeader(request *http.Request) {
	request.Header.Add(ContentTypeHeader, OctetStreamHeaderContentType)
}

func ExecuteRequestAndReadBodyAsString(c *http.Client, r *http.Request) (*[]byte, error) {
	response, err := c.Do(r)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return &body, nil
}
