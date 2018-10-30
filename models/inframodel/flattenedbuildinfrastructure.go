package inframodel

type FlattenedBuildInfrastructure struct {
	ServerName          string
	ServerGroup         string
	DeploymentLocations []string
	AppPoolNames        []string
	ServiceNames        []string
	TaskNames           []string
}
