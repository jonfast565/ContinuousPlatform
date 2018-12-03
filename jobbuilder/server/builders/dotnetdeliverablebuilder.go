package builders

import (
	"../../../clients/msbuildclient"
	"../../../clients/repoclient"
	"../../../fileutil"
	"../../../logging"
	"../../../models/filesysmodel"
	"../../../models/projectmodel"
	"../../../models/repomodel"
	"../../../stringutil"
	"github.com/ahmetb/go-linq"
)

var ValidProjectExtensions = []string{
	`^.*\.csproj$`,
	`^.*\.fsproj$`,
	`^.*\.vbproj$`,
}

var ValidSolutionExtensions = []string{
	`^.*\.sln$`,
}

var ValidPublishProfileExtensions = []string{
	`^.*\.pubxml$`,
}

type DotNetDeliverableBuildContext struct {
	repoClient          repoclient.RepoClient
	msBuildClient       msbuildclient.MsBuildClient
	solutionPaths       []string
	solutions           []*projectmodel.MsBuildSolution
	projectPaths        []string
	projects            []*projectmodel.MsBuildProject
	publishProfilePaths []string
	publishProfiles     []*projectmodel.MsBuildPublishProfile
	fileGraph           fileutil.FileGraph
	metadata            repomodel.RepositoryMetadata
}

func NewDotNetDeliverableBuildContext(metadata repomodel.RepositoryMetadata) *DotNetDeliverableBuildContext {
	graph := metadata.BuildGraph()
	context := DotNetDeliverableBuildContext{
		msBuildClient: msbuildclient.NewMsBuildClient(),
		repoClient:    repoclient.NewRepoClient(),
		fileGraph:     *graph,
		metadata:      metadata,
	}
	return &context
}

func (dndbc *DotNetDeliverableBuildContext) extractDeliverables() []*projectmodel.DotNetDeliverable {
	results := make([]*projectmodel.DotNetDeliverable, 0)
	for _, solution := range dndbc.solutions {
		results = append(results, &projectmodel.DotNetDeliverable{
			Repository:    dndbc.metadata.Name,
			RepositoryUrl: dndbc.metadata.OptionalUrl,
			Branch:        dndbc.metadata.Branch,
			Solution:      solution.Export(),
		})
	}
	return results
}

func (dndbc *DotNetDeliverableBuildContext) BuildContext() ([]*projectmodel.DotNetDeliverable, error) {
	err := dndbc.populatePaths()
	if err != nil {
		return nil, err
	}
	dndbc.populateSolutions()
	dndbc.populateProjects()
	dndbc.populatePublishProfiles()
	dndbc.resolveProjectReferencePaths()
	dndbc.resolveSolutionReferencePaths()
	for _, project := range dndbc.projects {
		dndbc.resolveProjectDependencies(project)
	}
	for _, solution := range dndbc.solutions {
		dndbc.linkProjectSolutions(solution)
	}
	result := dndbc.extractDeliverables()
	return result, nil
}

func (dndbc DotNetDeliverableBuildContext) GetContext() string {
	return "[" + dndbc.metadata.Name + " b. " + dndbc.metadata.Branch + "] "
}

func (dndbc DotNetDeliverableBuildContext) LogInfoWithContext(logLine string) {
	logging.LogInfo(dndbc.GetContext() + logLine)
}

func (dndbc DotNetDeliverableBuildContext) LogInfoMultilineWithContext(logLines ...string) {
	contextLines := []string{"Context: " + dndbc.GetContext()}
	contextLines = append(contextLines, logLines...)
	logging.LogInfoMultiline(contextLines...)
}

func (dndbc DotNetDeliverableBuildContext) LogErrorWithContext(err error) {
	logging.LogInfo(dndbc.GetContext() + err.Error())
}

func (dndbc *DotNetDeliverableBuildContext) populatePaths() error {
	solutionPaths, err := getSolutionPaths(dndbc.metadata)
	if err != nil {
		return err
	}

	if len(solutionPaths) == 0 {
		dndbc.LogInfoWithContext("Repo contains no solution files.")
		return nil
	}

	projectPaths, err := getProjectPaths(dndbc.metadata)
	if err != nil {
		return err
	}

	if len(projectPaths) == 0 {
		dndbc.LogInfoWithContext("Repo contains no project files.")
		return nil
	}

	publishProfilePaths, err := getPublishProfilePaths(dndbc.metadata)
	if err != nil {
		return err
	}

	dndbc.solutionPaths = solutionPaths
	dndbc.projectPaths = projectPaths
	dndbc.publishProfilePaths = publishProfilePaths

	return nil
}

func (dndbc *DotNetDeliverableBuildContext) populateSolutions() {
	var solutions []*projectmodel.MsBuildSolution
	for _, solutionPath := range dndbc.solutionPaths {
		dndbc.LogInfoWithContext("Downloading solution: " + solutionPath)

		solution := dndbc.getSolutionFromSourceControl(solutionPath)
		solution.AbsolutePath = solutionPath
		solutionFolderPath, err := dndbc.fileGraph.GetParentPath(solution.AbsolutePath)
		if err != nil {
			dndbc.LogInfoMultilineWithContext("Failed to find solution parent folder",
				"What: "+err.Error(),
				"Solution Name: "+solution.Name)
		}
		solution.FolderPath = *solutionFolderPath

		// filter the unvalidated paths by project extension, to avoid solution folders
		var relativePathsWithoutFolders []string
		relativePathsWithoutFoldersQuery := linq.
			From(solution.RelativeProjectPaths).
			WhereT(func(path string) bool {
				match, err := stringutil.StringMatchesOneOfRegStr(path, ValidProjectExtensions)
				if err != nil {
					return false
				}
				return match
			})
		relativePathsWithoutFoldersQuery.ToSlice(&relativePathsWithoutFolders)

		prettyPaths, err := dndbc.fileGraph.ValidatePathsFromRoot(relativePathsWithoutFolders, true)
		if err != nil {
			dndbc.LogInfoMultilineWithContext("Failed to find a solution project",
				"What: "+err.Error(),
				"Solution Name: "+solution.Name)
		}

		solution.AbsoluteProjectPaths = prettyPaths
		solutions = append(solutions, solution)
	}

	dndbc.solutions = solutions
}

func (dndbc *DotNetDeliverableBuildContext) populateProjects() {
	var projects []*projectmodel.MsBuildProject
	for _, projectPath := range dndbc.projectPaths {
		dndbc.LogInfoWithContext("Downloading project: " + projectPath)

		project := dndbc.getProjectFromSourceControl(projectPath)
		project.AbsolutePath = projectPath
		projectFolderPath, err := dndbc.fileGraph.GetParentPath(project.AbsolutePath)
		if err != nil {
			dndbc.LogInfoMultilineWithContext("Failed to find project parent folder",
				"What: "+err.Error(),
				"Project Name: "+project.Name)
		}

		project.FolderPath = *projectFolderPath
		projects = append(projects, project)
	}

	dndbc.projects = projects
}

func (dndbc *DotNetDeliverableBuildContext) populatePublishProfiles() {
	var publishProfiles []*projectmodel.MsBuildPublishProfile
	for _, publishProfilePath := range dndbc.publishProfilePaths {
		dndbc.LogInfoWithContext("Downloading publish profile: " + publishProfilePath)
		publishProfile := dndbc.getPublishProfileFromSourceControl(publishProfilePath)
		publishProfiles = append(publishProfiles, publishProfile)
	}
	dndbc.publishProfiles = publishProfiles
}

func BuildDotNetDeliverables(metadata repomodel.RepositoryMetadata) ([]*projectmodel.DotNetDeliverable, error) {
	results := make([]*projectmodel.DotNetDeliverable, 0)

	dndbc := NewDotNetDeliverableBuildContext(metadata)
	deliverables, err := dndbc.BuildContext()
	if err != nil {
		dndbc.LogInfoWithContext("Failed building .NET deliverables")
	}

	results = append(results, deliverables...)
	dndbc.LogInfoWithContext("Done building .NET deliverables")

	return results, nil
}

func (dndbc *DotNetDeliverableBuildContext) resolveSolutionReferencePaths() {
	for _, solution := range dndbc.solutions {

		solution.AbsoluteProjectPaths = make([]string, 0)
		for _, relativeProjectPath := range solution.RelativeProjectPaths {

			rootItem, err := dndbc.fileGraph.GetItemByRelativePath(solution.FolderPath, relativeProjectPath)
			if err != nil {
				dndbc.LogInfoMultilineWithContext("Could not find project at path.",
					"Path: "+relativeProjectPath,
					"Error: "+err.Error())
				continue
			}

			rootPath := (*rootItem).GetPathString()
			solution.AbsoluteProjectPaths = append(solution.AbsoluteProjectPaths, rootPath)
			found := false

			for _, project := range dndbc.projects {
				if rootPath != project.AbsolutePath {
					continue
				}
				solution.Projects = append(solution.Projects, project)
				found = true
				break
			}

			if !found {
				dndbc.LogInfoMultilineWithContext("Failed to find solution reference",
					"Project Path: "+rootPath,
					"Solution Name: "+solution.Name)
			}
		}
	}
}

func (dndbc *DotNetDeliverableBuildContext) resolveProjectDependencies(project *projectmodel.MsBuildProject) {
	dndbc.LogInfoWithContext("Resolving dependencies for: " + project.Name)
	project.ProjectDependencies = make([]*projectmodel.MsBuildProjectReference, 0)

	for _, absolutePath := range project.AbsoluteProjectReferencePaths {
		found := false
		for _, referenceProject := range dndbc.projects {
			if absolutePath != referenceProject.AbsolutePath ||
				absolutePath == project.AbsolutePath {
				continue
			}

			project.ProjectDependencies = append(project.ProjectDependencies, referenceProject.GetReference())
			found = true
			break
		}

		if !found {
			dndbc.LogInfoMultilineWithContext("Failed to find project reference",
				"Absolute Path: "+absolutePath,
				"Project Name: "+project.Name)
		}
	}
}

func (dndbc *DotNetDeliverableBuildContext) linkProjectSolutions(solution *projectmodel.MsBuildSolution) {
	for _, projectPath := range solution.AbsoluteProjectPaths {
		found := false
		for _, project := range dndbc.projects {
			if projectPath != project.AbsolutePath {
				continue
			}
			solutionReference := solution.GetSolutionReference()
			project.SolutionParents = append(project.SolutionParents, &solutionReference)
			// TODO: Understand this duplication, it's causing issues
			// solution.Projects = append(solution.Projects, project)
			found = true
			dndbc.LogInfoWithContext("Linked " + solution.Name + " -> " + project.Name)
			break
		}

		if !found {
			dndbc.LogInfoMultilineWithContext("Failed to link solution to project. Project with path not found.",
				"Solution Name: "+solution.Name,
				"Project Path: "+projectPath)
		}
	}
}

func (dndbc *DotNetDeliverableBuildContext) resolveProjectReferencePaths() {
	for _, project := range dndbc.projects {
		project.AbsoluteProjectReferencePaths = make([]string, 0)
		for _, relativeProjectPath := range project.RelativeProjectReferencePaths {
			fileGraphItem, err := dndbc.fileGraph.GetItemByRelativePath(project.FolderPath, relativeProjectPath)
			if err != nil {
				dndbc.LogInfoMultilineWithContext("Could not find relative project path. Ignoring.",
					"Project Name: "+project.Name,
					"Path: "+relativeProjectPath,
					"Error: "+err.Error())
				continue
			} else {
				resolvedProjectPath := (*fileGraphItem).GetPathString()
				project.AbsoluteProjectReferencePaths = append(project.AbsoluteProjectReferencePaths, resolvedProjectPath)
			}
		}
	}
}

func (dndbc *DotNetDeliverableBuildContext) getPublishProfileFromSourceControl(
	path string) *projectmodel.MsBuildPublishProfile {
	repoMetadata := getRepositoryFileMetadataFromPath(dndbc.metadata, path)

	payload, err := dndbc.repoClient.GetFile(repoMetadata)
	if err != nil {
		panic(err)
	}

	publishProfile, err := dndbc.msBuildClient.GetPublishProfile(*payload)
	if err != nil {
		panic(err)
	}

	return publishProfile
}

func (dndbc *DotNetDeliverableBuildContext) getProjectFromSourceControl(
	path string) *projectmodel.MsBuildProject {
	repoMetadata := getRepositoryFileMetadataFromPath(dndbc.metadata, path)

	payload, err := dndbc.repoClient.GetFile(repoMetadata)
	if err != nil {
		panic(err)
	}

	project, err := dndbc.msBuildClient.GetProject(*payload)
	if err != nil {
		panic(err)
	}

	return project
}

func (dndbc *DotNetDeliverableBuildContext) getSolutionFromSourceControl(
	path string) *projectmodel.MsBuildSolution {
	repoMetadata := getRepositoryFileMetadataFromPath(dndbc.metadata, path)

	payload, err := dndbc.repoClient.GetFile(repoMetadata)
	if err != nil {
		panic(err)
	}

	solution, err := dndbc.msBuildClient.GetSolution(*payload)
	if err != nil {
		panic(err)
	}

	return solution
}

func getRepositoryFileMetadataFromPath(metadata repomodel.RepositoryMetadata,
	path string) repomodel.RepositoryFileMetadata {
	fileMetadata := filesysmodel.FileSystemMetadata{
		Path: path,
		Type: filesysmodel.FileType,
	}
	return repomodel.NewRepositoryFileMetadata(metadata.Name, metadata.Branch, path, fileMetadata)
}

func getSolutionPaths(metadata repomodel.RepositoryMetadata) ([]string, error) {
	return metadata.GetMatchingFiles(ValidSolutionExtensions)
}

func getProjectPaths(metadata repomodel.RepositoryMetadata) ([]string, error) {
	return metadata.GetMatchingFiles(ValidProjectExtensions)
}

func getPublishProfilePaths(metadata repomodel.RepositoryMetadata) ([]string, error) {
	return metadata.GetMatchingFiles(ValidPublishProfileExtensions)
}
