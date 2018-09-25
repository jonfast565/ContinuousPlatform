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
	Name      string
	Url       string
	Jobs      []JobMetadata
	JobRecord JobRecord
}

type JobRecord struct {
	TopLevelFolder    string
	MidLevel1Folder   string
	MidLevel2Folder   string
	BottomLevelFolder string
}

type Edit struct {
	Name      string
	Url       string
	EditType  EditType
	JobRecord JobRecord
}

type NewFolderRequest struct {
	Name   string
	Mode   string
	From   string
	Submit string
}

func CreateFolderRequest(name string) NewFolderRequest {
	return NewFolderRequest{
		Name:   name,
		Mode:   "com.cloudbees.hudson.plugins.folder.Folder",
		From:   "",
		Submit: "OK",
	}
}
