package jenkinsmodel

import (
	"../../stringutil"
	"sort"
)

type JenkinsJobMetadata struct {
	Name  string
	Url   string
	Jobs  []JenkinsJobMetadata
	Class string `json:"_class"`
}

type JenkinsJobKey struct {
	Keys []string
	Type string
}

type JenkinsJobKeyList []JenkinsJobKey

func (jjkl JenkinsJobKeyList) Len() int {
	return len(jjkl)
}
func (jjkl JenkinsJobKeyList) Swap(i, j int) {
	jjkl[i], jjkl[j] = jjkl[j], jjkl[i]
}
func (jjkl JenkinsJobKeyList) Less(i, j int) bool {
	return len(jjkl[i].Keys) < len(jjkl[j].Keys)
}

func (jjkl JenkinsJobKeyList) KeyAlreadyExists(keys []string) bool {
	for _, jobKeys := range jjkl {
		if stringutil.StringArrayCompare(jobKeys.Keys, keys) {
			return true
		}
	}
	return false
}

func (jjm JenkinsJobMetadata) GetFlattenedKeys() *JenkinsJobKeyList {
	result := getFlattenedKeysInternal(nil, jjm)
	sort.Sort(result)
	return &result
}

func getFlattenedKeysInternal(currentKey *JenkinsJobKey, metadata JenkinsJobMetadata) JenkinsJobKeyList {
	var result JenkinsJobKeyList
	var newKey JenkinsJobKey

	if currentKey == nil {
		newKey = JenkinsJobKey{
			Keys: []string{metadata.Name},
			Type: metadata.Class,
		}
	} else {
		newKey = JenkinsJobKey{
			Keys: append(currentKey.Keys, metadata.Name),
			Type: metadata.Class,
		}
	}

	result = append(result, newKey)
	for _, record := range metadata.Jobs {
		internalKeys := getFlattenedKeysInternal(&newKey, record)
		result = append(result, internalKeys...)
	}
	return result
}
