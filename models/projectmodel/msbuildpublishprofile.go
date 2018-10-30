package projectmodel

type MsBuildPublishProfile struct {
	PublishUrl string
	Name       string
	Failed     bool
	Exception  DotNetException
}
