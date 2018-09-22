package teamservices

type TeamServicesGitFileList struct {
	Value []TeamServicesGitFileModel
	Count int
}

type TeamServicesGitFileModel struct {
	ObjectId      string
	GitObjectType string
	CommitId      string
	Path          string
	Url           string
}

type TeamServicesGitRefsList struct {
	Value []TeamServicesGitRefsModel
	Count int
}

type TeamServicesGitRefsModel struct {
	Name     string
	ObjectId string
	Url      string
}

type TeamServicesGitRepositoryList struct {
	Value []TeamServicesGitRepositoryModel
	Count int
}

type TeamServicesGitRepositoryModel struct {
	Id        string
	Name      string
	Url       string
	Project   TeamServicesGitRepositoryProjectModel
	RemoteUrl string
}

type TeamServicesGitRepositoryProjectModel struct {
	Id    string
	Name  string
	Url   string
	State string
}
