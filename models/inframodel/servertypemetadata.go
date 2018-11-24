package inframodel

type ServerType struct {
	ServerName string
	ServerType string
}

type ServerTypeMetadata struct {
	ServerName          string
	ServerType          string
	EnvironmentName     string
	DeploymentLocations []string
	AppPoolNames        []string
	ServiceNames        []string
	TaskNames           []string
}
