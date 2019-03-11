package jenkinsmodel

import (
	"github.com/jonfast565/continuous-platform/constants"
	"github.com/jonfast565/continuous-platform/stringutil"
)

type JenkinsJobKeyList []JenkinsJobKey

func (jjkl JenkinsJobKeyList) Len() int {
	return len(jjkl)
}
func (jjkl JenkinsJobKeyList) Swap(i, j int) {
	jjkl[i], jjkl[j] = jjkl[j], jjkl[i]
}
func (jjkl JenkinsJobKeyList) Less(i, j int) bool {
	return len(jjkl[i].Keys) < len(jjkl[j].Keys) &&
		stringutil.StringArrayCompareNumeric(jjkl[i].Keys, jjkl[j].Keys) < 0
}

func (jjkl JenkinsJobKeyList) SanitizeKeyList() {
	for _, kl := range jjkl {
		kl.SanitizeKeys()
	}
}

func (jjkl JenkinsJobKeyList) KeyAlreadyExists(keys []string) bool {
	for _, jobKeys := range jjkl {
		if stringutil.StringArrayCompare(jobKeys.Keys, keys) {
			return true
		}
	}
	return false
}

func (jjkl JenkinsJobKeyList) PartialKeyAlreadyExists(keys []string) bool {
	for _, jobKeys := range jjkl {
		if stringutil.StringArrayContainsArray(jobKeys.Keys, keys) {
			return true
		}
	}
	return false
}

func (jjkl JenkinsJobKeyList) CleanRawBuildServerKeys() *JenkinsJobKeyList {
	var cleanedKeys JenkinsJobKeyList
	for _, key := range jjkl {
		if len(key.Keys) == 0 {
			continue
		}
		if key.Type == BuildServer {
			continue
		}
		newKeyList := key.Keys[1:len(key.Keys)]
		// TODO: This is never not true...
		if key.Keys[0] == constants.JenkinsRootVariable {
			cleanedKeys = append(cleanedKeys, JenkinsJobKey{
				Keys: newKeyList,
				Type: key.Type,
			})
		}
	}
	return &cleanedKeys
}
