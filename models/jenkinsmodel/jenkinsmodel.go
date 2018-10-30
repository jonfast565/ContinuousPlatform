package jenkinsmodel

const (
	JenkinsMaximumJobDepth int = 4
)

type JenkinsEditType int

const (
	RemoveJobFolder JenkinsEditType = 0
	AddUpdateJob    JenkinsEditType = 1
	AddFolder       JenkinsEditType = 2
)
