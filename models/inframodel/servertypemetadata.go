package inframodel

import (
	"../../stringutil"
	"sort"
)

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

type ServerTypeMetadataList []ServerTypeMetadata

func (stml ServerTypeMetadataList) GetEnvironments() []string {
	var result []string
	for _, environment := range stml {
		if !stringutil.StringArrayContains(result, environment.EnvironmentName) {
			result = append(result, environment.EnvironmentName)
		}
	}
	alpha := stringutil.AlphabeticArray(result)
	sort.Sort(alpha)
	return alpha
}
