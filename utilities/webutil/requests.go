package webutil

import (
	"bytes"
	"encoding/json"
	"github.com/go-errors/errors"
	"github.com/yosssi/gohtml"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	AuthorizationHeader              string = "Authorization"
	ContentTypeHeader                string = "Content-Type"
	ApplicationJsonHeaderContentType string = "application/json"
	OctetStreamHeaderContentType     string = "application/octet-stream"
	XmlHeaderContentType             string = "text/xml"
	FormUnEncodedHeaderContentType   string = "application/x-www-form-urlencoded"
	HtmlDocumentHeader               string = "<!DOCTYPE html>"
	NoDataError                      string = "Nothing Returned"
)

var windowsNewLines = "\r\n"
var windowsNewLinesByte = []byte{13, 10}
var unixNewLines = "\n"
var unixNewLinesByte = []byte{10}

// Adds a JSON header to an HTTP request
func AddJsonHeader(request *http.Request) {
	request.Header.Add(ContentTypeHeader, ApplicationJsonHeaderContentType)
}

// Adds an Octet/Binary header to an HTTP request
func AddOctetHeader(request *http.Request) {
	request.Header.Add(ContentTypeHeader, OctetStreamHeaderContentType)
}

// Adds an XML header to an HTTP request
func AddXmlHeader(request *http.Request) {
	request.Header.Add(ContentTypeHeader, XmlHeaderContentType)
}

// Adds an unencoded form (X-WWW-FORM-URLENCODED) header to an HTTP request
// In this case, form data must be passed in the URL using query parameters
func AddFormHeader(request *http.Request) {
	request.Header.Add(ContentTypeHeader, FormUnEncodedHeaderContentType)
}

// Adds a bearer token to an HTTP request
func AddBearerToken(request *http.Request, bearerToken string) {
	request.Header.Add(AuthorizationHeader, "Bearer "+bearerToken)
}

// Executes a request and populates its body via an interface
// Returns an error if serdes fails, or in the case of a bad status
func ExecuteRequestAndReadJsonBody(r *http.Request, object interface{}) error {
	c := NewHttpClient()
	response, err := c.Do(r)
	if err != nil {
		return err
	}
	if response.Body != nil {
		defer response.Body.Close()
	}

	if response.StatusCode >= 400 {
		return errors.New("bad status (" + response.Status + ") returned from server")
	}

	err = json.NewDecoder(response.Body).Decode(&object)
	if err != nil {
		return err
	}

	return nil
}

// Executes a request and returns a byte array with success or fail indicator
// Returns an error if serdes fails, or in the case of a bad status
func ExecuteRequestAndReadBinaryBody(r *http.Request) (*[]byte, error) {
	c := NewHttpClient()
	response, err := c.Do(r)
	if err != nil {
		return nil, err
	}
	if response.Body != nil {
		defer response.Body.Close()
	}

	resultBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		if err.Error() != "EOF" {
			return nil, err
		} else {
			resultBytes = []byte(NoDataError)
		}
	}

	if response.StatusCode >= 400 {
		resultStr := string(resultBytes)
		if strings.Contains(resultStr, HtmlDocumentHeader) {
			resultStr = gohtml.Format(resultStr)
		}
		return nil, errors.New("bad status (" + response.Status + ") returned from server: \n" + resultStr)
	}

	resultBytes = bytes.Replace(resultBytes, unixNewLinesByte, windowsNewLinesByte, -1)
	return &resultBytes, nil
}

// Executes a request and returns a string array with success or fail indicator
// Returns an error if serdes fails, or in the case of a bad status
func ExecuteRequestAndReadStringBody(r *http.Request, ignoreStatusCode bool) (*string, error) {
	c := NewHttpClient()
	response, err := c.Do(r)
	if err != nil {
		return nil, err
	}
	if response.Body != nil {
		defer response.Body.Close()
	}

	resultBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		if err.Error() != "EOF" {
			return nil, err
		} else {
			resultBytes = []byte(NoDataError)
		}
	}

	if response.StatusCode >= 400 && !ignoreStatusCode {
		resultStr := string(resultBytes)
		if strings.Contains(resultStr, HtmlDocumentHeader) {
			resultStr = gohtml.Format(resultStr)
		}
		return nil, errors.New("bad status (" + response.Status + ") returned from server: \n" + resultStr)
	}

	// build a string and return
	builder := strings.Builder{}
	_, err = builder.Write(resultBytes)
	if err != nil {
		panic(err)
	}
	s := builder.String()

	s = strings.Replace(s, windowsNewLines, unixNewLines, -1)
	return &s, nil
}

// Executes a request and return success/fail with no resulting payload
// Returns an error if serdes fails, or in the case of a bad status
func ExecuteRequestWithoutRead(r *http.Request) error {
	c := NewHttpClient()
	response, err := c.Do(r)
	if err != nil {
		return err
	}
	if response.Body != nil {
		defer response.Body.Close()
	}

	if response.StatusCode >= 400 {
		return errors.New("bad status (" + response.Status + ") returned from server, no response body")
	}

	_, err = io.Copy(ioutil.Discard, response.Body)
	if err != nil {
		return err
	}

	return nil
}
