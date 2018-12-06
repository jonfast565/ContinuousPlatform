package jenkinsmodel

type JenkinsKeyStack JenkinsJobKeyList

func ListToStack(list JenkinsJobKeyList) JenkinsKeyStack {
	return JenkinsKeyStack(list)
}

func (jks *JenkinsKeyStack) Push(key JenkinsJobKey) {
	*jks = append(*jks, key)
}

func (jks *JenkinsKeyStack) Pop() *JenkinsJobKey {
	lastItem := (*jks)[len(*jks)-1]
	*jks = (*jks)[0 : len(*jks)-1]
	return &lastItem
}

func (jks JenkinsKeyStack) ReadStackAsList() JenkinsJobKeyList {
	var result JenkinsJobKeyList
	for _, stackItem := range jks {
		result = append(result, stackItem)
	}
	return result
}
