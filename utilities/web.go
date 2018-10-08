package utilities

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	AuthorizationHeader              string = "Authorization"
	ContentTypeHeader                string = "Content-Type"
	ApplicationJsonHeaderContentType string = "application/json"
	OctetStreamHeaderContentType     string = "application/octet-stream"
)

var windowsNewLines = "\r\n"
var windowsNewLinesByte = []byte{13, 10}
var unixNewLines = "\n"
var unixNewLinesByte = []byte{10}

func AddJsonHeader(request *http.Request) {
	request.Header.Add(ContentTypeHeader, ApplicationJsonHeaderContentType)
}

func AddOctetHeader(request *http.Request) {
	request.Header.Add(ContentTypeHeader, OctetStreamHeaderContentType)
}

func AddBearerToken(request *http.Request, bearerToken string) {
	request.Header.Add(AuthorizationHeader, "Bearer "+bearerToken)
}

func ExecuteRequestAndReadJsonBody(c *http.Client, r *http.Request, object interface{}) error {
	response, err := c.Do(r)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&object)
	if err != nil {
		return err
	}
	return nil
}

func ExecuteRequestAndReadBinaryBody(c *http.Client, r *http.Request) (*[]byte, error) {
	response, err := c.Do(r)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	resultBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	resultBytes = bytes.Replace(resultBytes, unixNewLinesByte, windowsNewLinesByte, -1)
	return &resultBytes, nil
}

func ExecuteRequestAndReadStringBody(c *http.Client, r *http.Request) (*string, error) {
	response, err := c.Do(r)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}
	resultBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// build a string and return
	// TODO: Ensure correctness
	builder := strings.Builder{}
	builder.Write(resultBytes)
	var s string
	builder.WriteString(s)

	s = strings.Replace(s, windowsNewLines, unixNewLines, -1)
	return &s, nil
}
