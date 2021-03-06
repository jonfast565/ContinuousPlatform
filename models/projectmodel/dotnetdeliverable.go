package projectmodel

import (
	"github.com/jonfast565/continuous-platform/models/genmodel"
	"github.com/jonfast565/continuous-platform/models/inframodel"
	"strings"
)

type DotNetDeliverable struct {
	Repository          string
	RepositoryUrl       string
	Branch              string
	Solution            *MsBuildSolutionExport
	DependencySolutions []*MsBuildSolutionReference
}

type FlattenedDotNetDeliverable struct {
	Repository          string
	RepositoryUrl       string
	Branch              string
	Solution            *MsBuildSolutionExport
	DependencySolutions []*MsBuildSolutionReference
	Project             *MsBuildProjectExport
}

func (dnd DotNetDeliverable) GetFlattenedDeliverables() *[]FlattenedDotNetDeliverable {
	var result []FlattenedDotNetDeliverable
	for _, project := range dnd.Solution.Projects {
		result = append(result, FlattenedDotNetDeliverable{
			Repository:          dnd.Repository,
			RepositoryUrl:       dnd.RepositoryUrl,
			Branch:              dnd.Branch,
			Solution:            dnd.Solution,
			DependencySolutions: dnd.DependencySolutions,
			Project:             project,
		})
	}
	return &result
}

func (fdnd FlattenedDotNetDeliverable) GetScriptKey(template genmodel.ScriptTemplate) []string {
	return []string{
		fdnd.Repository,
		fdnd.Branch,
		fdnd.Solution.Name,
		fdnd.Project.Name + " - " + string(template.Type),
	}
}

func (fdnd FlattenedDotNetDeliverable) GetScriptKeyString(template genmodel.ScriptTemplate) string {
	keyArray := fdnd.GetScriptKey(template)
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
