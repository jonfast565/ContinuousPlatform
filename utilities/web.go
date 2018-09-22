package utilities

import (
	"encoding/json"
	"net/http"
)

const (
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

func ExecuteRequestAndReadJsonBody(c *http.Client, r *http.Request, object interface{}) error {
	response, err := c.Do(r)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	err := json.NewDecoder(response.Body).Decode(&object)
	if err != nil {
		return err
	}
}
