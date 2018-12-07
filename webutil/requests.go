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

func AddJsonHeader(request *http.Request) {
	request.Header.Add(ContentTypeHeader, ApplicationJsonHeaderContentType)
}

func AddOctetHeader(request *http.Request) {
	request.Header.Add(ContentTypeHeader, OctetStreamHeaderContentType)
}

func AddXmlHeader(request *http.Request) {
	request.Header.Add(ContentTypeHeader, XmlHeaderContentType)
}

func AddFormHeader(request *http.Request) {
	request.Header.Add(ContentTypeHeader, FormUnEncodedHeaderContentType)
}

func AddBearerToken(request *http.Request, bearerToken string) {
	request.Header.Add(AuthorizationHeader, "Bearer "+bearerToken)
}

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

func ExecuteRequestAndReadStringBody(r *http.Request) (*string, error) {
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

	// build a string and return
	builder := strings.Builder{}
	builder.Write(resultBytes)
	var s string
	builder.WriteString(s)

	s = strings.Replace(s, windowsNewLines, unixNewLines, -1)
	return &s, nil
}

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
