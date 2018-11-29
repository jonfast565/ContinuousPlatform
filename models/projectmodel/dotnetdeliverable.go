package projectmodel

import (
	"../inframodel"
	"strings"
)

type DotNetDeliverable struct {
	Repository string
	Branch     string
	Solution   *MsBuildSolutionExport
}

type FlattenedDotNetDeliverable struct {
	Repository string
	Branch     string
	Solution   *MsBuildSolutionExport
	Project    *MsBuildProjectExport
}

func (dnd DotNetDeliverable) GetFlattenedDeliverables() *[]FlattenedDotNetDeliverable {
	var result []FlattenedDotNetDeliverable
	for _, project := range dnd.Solution.Projects {
		result = append(result, FlattenedDotNetDeliverable{
			Repository: dnd.Repository,
			Branch:     dnd.Branch,
			Solution:   dnd.Solution,
			Project:    project,
		})
	}
	return &result
}

func (fdnd FlattenedDotNetDeliverable) GetScriptKey() []string {
	return []string{fdnd.Repository, fdnd.Branch, fdnd.Solution.Name, fdnd.Project.Name}
}

func (fdnd FlattenedDotNetDeliverable) GetScriptKeyString() string {
	keyArray := fdnd.GetScriptKey()
	keyString := strings.Join(keyArray, "-")
	return keyString
}

func (fdnd FlattenedDotNetDeliverable) GetRepositoryKey() inframodel.ResourceKey {
	return inframodel.ResourceKey{
		RepositoryName: fdnd.Repository,
		SolutionName:   fdnd.Solution.Name,
		ProjectName:    fdnd.Project.Name,
	}
}
