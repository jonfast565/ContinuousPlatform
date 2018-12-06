package jenkinsmodel

import "strings"

type JenkinsJobKey struct {
	Keys []string
	Type JenkinsJobType
}

func (jjk JenkinsJobKey) SanitizeKeys() {
	for i, key := range jjk.Keys {
		jjk.Keys[i] = strings.Replace(key, "/", "-", -1)
	}
}
