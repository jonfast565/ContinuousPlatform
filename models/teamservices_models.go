package models

type TsGitFileList struct {
	Value []TsGitFileModel
	Count int
}

type TsGitFileModel struct {
	ObjectId      string
	GitObjectType string
	CommitId      string
	Path          string
	Url           string
}

type TsGitRefsList struct {
	Value []TsGitRefsModel
	Count int
}

type TsGitRefsModel struct {
	Name     string
	ObjectId string
	Url      string
}

type TsGitRepositoryList struct {
	Value []TsGitRepositoryModel
	Count int
}

type TsGitRepositoryModel struct {
	Id        string
	Name      string
	Url       string
	Project   TsGitRepositoryProjectModel
	RemoteUrl string
}

type TsGitRepositoryProjectModel struct {
	Id    string
	Name  string
	Url   string
	State string
}
