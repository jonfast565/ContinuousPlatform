package webutil

import (
	"net/url"
)

type MyUrl struct {
	Scheme        string
	Host          string
	Port          string
	PathFragments []string
	QueryValues   map[string]string
}

func NewEmptyUrl() MyUrl {
	return MyUrl{
		Scheme:        "",
		Host:          "",
		Port:          "",
		PathFragments: make([]string, 0),
		QueryValues:   make(map[string]string, 0),
	}
}

func (u *MyUrl) SetBase(scheme string, host string, port string) {
	u.Scheme = scheme
	u.Host = host
	u.Port = port
}

func (u *MyUrl) AppendPathFragment(fragment string) {
	u.PathFragments = append(u.PathFragments, fragment)
}

func (u *MyUrl) AppendPathFragments(fragment []string) {
	u.PathFragments = append(u.PathFragments, fragment...)
}

func (u *MyUrl) AppendQueryValue(key string, value string) {
	u.QueryValues[key] = value
}

func (u *MyUrl) GetQueryValue() string {
	var queryString string
	queryLength := len(u.QueryValues)
	counter := 0
	for key, value := range u.QueryValues {
		queryString += url.QueryEscape(key) +
			"=" + url.QueryEscape(value)
		if counter != queryLength-1 {
			queryString += "&"
		}
		counter++
	}
	result := queryString
	return result
}

func (u *MyUrl) GetFragmentValue() string {
	var pathFragment string
	for i, fragment := range u.PathFragments {
		if i == len(u.PathFragments)-1 {
			pathFragment += url.PathEscape(fragment)
		} else {
			pathFragment += url.PathEscape(fragment) + "/"
		}
	}
	result := pathFragment
	return result
}

func (u *MyUrl) GetBasePath() string {
	result := u.Scheme + "://" + u.Host
	if u.Port != "" {
		result += ":" + u.Port
	}
	result += "/"
	return result
}

func (u *MyUrl) GetUrlStringValue() string {
	result := u.GetBasePath() + u.GetFragmentValue()
	if len(u.QueryValues) > 0 {
		result += "?" + u.GetQueryValue()
	}
	return result
}
