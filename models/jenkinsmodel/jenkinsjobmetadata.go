package jenkinsmodel

import (
	"github.com/jonfast565/continuous-platform/logging"
	"net/url"
	"sort"
)

// import "sort"

type JenkinsJobMetadata struct {
	Name  string
	Url   string
	Jobs  []JenkinsJobMetadata
	Class JenkinsJobType `json:"_class"`
}

func (jjm JenkinsJobMetadata) GetFlattenedKeys() *JenkinsJobKeyList {
	var jenkinsKeyStack JenkinsJobMetadataStack
	keyChan := make(chan JenkinsJobKey)

	getFlattenedKeysInternal(jjm, jenkinsKeyStack, keyChan)
	var resultKeys JenkinsJobKeyList
	for {
		noMore := false
		select {
		case msg := <-keyChan:
			if !resultKeys.KeyAlreadyExists(msg.Keys) {
				resultKeys = append(resultKeys, msg)
			}
		default:
			logging.LogInfo("No more keys received")
			noMore = true
		}
		if noMore {
			break
		}
	}

	// key cleanup & sort
	cleanedKeys := resultKeys.CleanRawBuildServerKeys()
	sort.Sort(cleanedKeys)
	return cleanedKeys
}

func getFlattenedKeysInternal(
	currentMetadata JenkinsJobMetadata,
	currentStack JenkinsJobMetadataStack,
	keyChan chan JenkinsJobKey) {

	currentStack.Push(currentMetadata)
	metadataList := currentStack.ReadStackAsList()

	var jobKey JenkinsJobKey
	keyStrings := metadataList.GetKeyNames()
	jobKey.Keys = keyStrings

	if len(currentMetadata.Jobs) > 0 {
		jobKey.Type = Folder
		go func() { keyChan <- jobKey }()
		for _, childMetadata := range currentMetadata.Jobs {
			getFlattenedKeysInternal(childMetadata, currentStack, keyChan)
		}
	} else {
		jobKey.Type = currentMetadata.Class
		go func() { keyChan <- jobKey }()
	}
	currentStack.Pop()
}

type JenkinsJobMetadataList []JenkinsJobMetadata
type JenkinsJobMetadataStack JenkinsJobMetadataList

func (jjml JenkinsJobMetadataList) GetKeyNames() []string {
	var result []string
	for _, metadataItem := range jjml {
		// TODO: Need to know if this is necessary?
		unescapedName, err := url.PathUnescape(metadataItem.Name)
		if err != nil {
			panic(err)
		}
		result = append(result, unescapedName)
	}
	return result
}

func (jjms *JenkinsJobMetadataStack) Push(key JenkinsJobMetadata) {
	*jjms = append(*jjms, key)
}

func (jjms *JenkinsJobMetadataStack) Top() *JenkinsJobMetadata {
	lastItem := (*jjms)[len(*jjms)-1]
	return &lastItem
}

func (jjms *JenkinsJobMetadataStack) Pop() *JenkinsJobMetadata {
	lastItem := (*jjms)[len(*jjms)-1]
	*jjms = (*jjms)[0 : len(*jjms)-1]
	return &lastItem
}

func (jjms JenkinsJobMetadataStack) ReadStackAsList() JenkinsJobMetadataList {
	var result JenkinsJobMetadataList
	for _, stackItem := range jjms {
		result = append(result, stackItem)
	}
	return result
}
