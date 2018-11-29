package jenkinsmodel

type JenkinsJobMetadata struct {
	Name string
	Url  string
	Jobs []JenkinsJobMetadata
}

func (jjm JenkinsJobMetadata) GetFlattenedKeys() [][]string {
	result := getFlattenedKeysInternal([]string{}, jjm)
	return result
}

func getFlattenedKeysInternal(currentKeys []string, metadata JenkinsJobMetadata) [][]string {
	var result [][]string
	newKeys := append(currentKeys, metadata.Name)
	result = append(result, newKeys)
	for _, record := range metadata.Jobs {
		internalKeys := getFlattenedKeysInternal(newKeys, record)
		result = append(result, internalKeys...)
	}
	return result
}
