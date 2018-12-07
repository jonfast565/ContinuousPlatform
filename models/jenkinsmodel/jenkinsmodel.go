package jenkinsmodel

const (
	JenkinsMaximumJobDepth int = 4
)

type JenkinsEditType int

const (
	RemoveJobFolder JenkinsEditType = 0
	AddJob          JenkinsEditType = 1
	UpdateJob       JenkinsEditType = 2
	AddFolder       JenkinsEditType = 3
)
