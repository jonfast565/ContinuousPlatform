package jenkinsmodel

type JenkinsJobMetadata struct {
	Name  string
	Url   string
	Jobs  []JenkinsJobMetadata
	Class JenkinsJobType `json:"_class"`
}

func (jjm JenkinsJobMetadata) GetFlattenedKeys() *JenkinsJobKeyList {
	// TODO: Reimplement key flattening...... w/o recursion

	// key cleanup & sort
	//cleanedKeys := *cleanKeys()
	//sort.Sort(cleanedKeys)
	//return &cleanedKeys
	return nil
}
