package projectmodel

type MsBuildPublishProfile struct {
	PublishUrl   string
	Name         string
	AbsolutePath string
	FolderPath   string
	Failed       bool
	Exception    DotNetException
}
