package jenkins

const (
	MaximumJobDepth int = 4
)

type EditType int

const (
	RemoveJobFolder EditType = 0
	AddUpdateJob    EditType = 1
	AddFolder       EditType = 2
)

type Crumb struct {
	Crumb             string
	CrumbRequestField string
}

type JobMetadata struct {
	Name string
	Url  string
	Jobs []JobMetadata
}

type JobRecord struct {
	TopLevelFolder    string
	MidLevel1Folder   string
	MidLevel2Folder   string
	BottomLevelFolder string
}

type JobRequest struct {
	FolderUrl string
	Record    JobRecord
}

type Edit struct {
	Name      string
	Url       string
	EditType  EditType
	JobRecord JobRecord
}
