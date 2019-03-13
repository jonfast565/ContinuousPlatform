package generators

import (
	"github.com/ahmetb/go-linq"
	"github.com/jonfast565/continuous-platform/clients/persistenceclient"
	"github.com/jonfast565/continuous-platform/models/genmodel"
	"github.com/jonfast565/continuous-platform/models/inframodel"
	"github.com/jonfast565/continuous-platform/models/jobmodel"
	"github.com/jonfast565/continuous-platform/models/projectmodel"
	"github.com/jonfast565/continuous-platform/utilities/logging"
	"github.com/jonfast565/continuous-platform/utilities/timeutil"
	"github.com/satori/go.uuid"
)

// Dotnet script generator object
// Contains all generated scripts
type DotNetScriptGenerator struct {
	persistenceClient                persistenceclient.PersistenceClient
	BuildScripts                     []genmodel.ScriptTemplate
	DeployScripts                    []genmodel.ScriptTemplate
	BuildDeployScripts               []genmodel.ScriptTemplate
	EnvironmentInfrastructureScripts []genmodel.ScriptTemplate
	ResourceList                     inframodel.ResourceList
}

// Script generator constructor
func NewDotNetScriptGenerator() *DotNetScriptGenerator {
	templateList := genmodel.NewScriptTemplateList()

	scriptGenerator := DotNetScriptGenerator{
		persistenceClient:                persistenceclient.NewPersistenceClient(),
		BuildScripts:                     []genmodel.ScriptTemplate{},
		DeployScripts:                    []genmodel.ScriptTemplate{},
		BuildDeployScripts:               []genmodel.ScriptTemplate{},
		EnvironmentInfrastructureScripts: []genmodel.ScriptTemplate{},
		ResourceList:                     inframodel.ResourceList{},
	}

	for _, template := range templateList.Templates {
		if !template.Enabled {
			continue
		}
		if template.Framework != genmodel.DotNet {
			continue
		}
		switch template.Type {
		case genmodel.Build:
			scriptGenerator.BuildScripts = append(scriptGenerator.BuildScripts, template)
			break
		case genmodel.Deploy:
			scriptGenerator.DeployScripts = append(scriptGenerator.DeployScripts, template)
			break
		case genmodel.BuildDeploy:
			scriptGenerator.BuildDeployScripts = append(scriptGenerator.BuildDeployScripts, template)
			break
		case genmodel.EnvironmentInfrastructure:
			scriptGenerator.EnvironmentInfrastructureScripts = append(
				scriptGenerator.EnvironmentInfrastructureScripts, template)
			break
		default:
			break
		}
	}

	resourceList, err := scriptGenerator.persistenceClient.GetResourceList()
	if err != nil {
		panic(err)
	}

	scriptGenerator.ResourceList = *resourceList
	return &scriptGenerator
}

type DotNetBuildScriptHeader struct {
	Deliverable            projectmodel.FlattenedDotNetDeliverable
	Solution               projectmodel.MsBuildSolutionExport
	Solutions              []*projectmodel.MsBuildSolutionReference
	Project                projectmodel.MsBuildProjectExport
	ProjectFolder          string
	PublishProfiles        []*projectmodel.MsBuildPublishProfile
	TargetFrameworks       []string
	DefaultNamespace       string
	SolutionConfigurations []string
	CanonicalId            string
	DashedCanonicalId      string
	Hash                   string
	GeneratedDateTime      string
}

func NewDotNetBuildScriptHeader(
	dnd projectmodel.FlattenedDotNetDeliverable,
	template genmodel.ScriptTemplate) *DotNetBuildScriptHeader {
	uid, _ := uuid.NewV4()
	return &DotNetBuildScriptHeader{
		Deliverable:            dnd,
		Solution:               *dnd.Solution,
		Solutions:              dnd.DependencySolutions,
		Project:                *dnd.Project,
		ProjectFolder:          dnd.Project.FolderPath,
		PublishProfiles:        dnd.Project.PublishProfiles,
		TargetFrameworks:       dnd.Project.TargetFrameworks,
		DefaultNamespace:       dnd.Project.DefaultNamespace,
		SolutionConfigurations: dnd.Solution.Configurations,
		CanonicalId:            dnd.GetScriptKeyString(template),
		DashedCanonicalId:      dnd.GetScriptKeyString(template),
		Hash:                   uid.String(),
		GeneratedDateTime:      timeutil.GetCurrentTime(),
	}
}

type DotNetDeployScriptHeader struct {
	Deliverable       projectmodel.FlattenedDotNetDeliverable
	Solution          projectmodel.MsBuildSolutionExport
	Project           projectmodel.MsBuildProjectExport
	Infrastructure    inframodel.ServerTypeMetadataList
	Environments      []string
	CanonicalId       string
	DashedCanonicalId string
	Hash              string
	GeneratedDateTime string
}

func NewDotNetDeployScriptHeader(
	dnd projectmodel.FlattenedDotNetDeliverable,
	template genmodel.ScriptTemplate,
	bim *inframodel.BuildInfrastructureMetadata) *DotNetDeployScriptHeader {
	uid, _ := uuid.NewV4()
	return &DotNetDeployScriptHeader{
		Deliverable:       dnd,
		Solution:          *dnd.Solution,
		Project:           *dnd.Project,
		Infrastructure:    bim.Metadata,
		Environments:      inframodel.ServerTypeMetadataList(bim.Metadata).GetEnvironments(),
		CanonicalId:       dnd.GetScriptKeyString(template),
		DashedCanonicalId: dnd.GetScriptKeyString(template),
		Hash:              uid.String(),
		GeneratedDateTime: timeutil.GetCurrentTime(),
	}
}

type DotNetBuildDeployScriptHeader struct {
	Deliverable            projectmodel.FlattenedDotNetDeliverable
	Solution               projectmodel.MsBuildSolutionExport
	Solutions              []*projectmodel.MsBuildSolutionReference
	Project                projectmodel.MsBuildProjectExport
	ProjectFolder          string
	PublishProfiles        []*projectmodel.MsBuildPublishProfile
	TargetFrameworks       []string
	DefaultNamespace       string
	SolutionConfigurations []string
	Infrastructure         inframodel.ServerTypeMetadataList
	Environments           []string
	CanonicalId            string
	DashedCanonicalId      string
	Hash                   string
	GeneratedDateTime      string
}

func NewDotNetBuildDeployScriptHeader(
	dnd projectmodel.FlattenedDotNetDeliverable,
	template genmodel.ScriptTemplate,
	bim *inframodel.BuildInfrastructureMetadata) *DotNetBuildDeployScriptHeader {
	uid, _ := uuid.NewV4()
	return &DotNetBuildDeployScriptHeader{
		Deliverable:            dnd,
		Solution:               *dnd.Solution,
		Solutions:              dnd.DependencySolutions,
		Project:                *dnd.Project,
		ProjectFolder:          dnd.Project.FolderPath,
		PublishProfiles:        dnd.Project.PublishProfiles,
		TargetFrameworks:       dnd.Project.TargetFrameworks,
		DefaultNamespace:       dnd.Project.DefaultNamespace,
		SolutionConfigurations: dnd.Solution.Configurations,
		Infrastructure:         bim.Metadata,
		Environments:           inframodel.ServerTypeMetadataList(bim.Metadata).GetEnvironments(),
		CanonicalId:            dnd.GetScriptKeyString(template),
		DashedCanonicalId:      dnd.GetScriptKeyString(template),
		Hash:                   uid.String(),
		GeneratedDateTime:      timeutil.GetCurrentTime(),
	}
}

type DotNetEnvironmentInfrastructureScriptHeader struct {
}

func NewDotNetEnvironmentInfrastructureScriptHeader() *DotNetEnvironmentInfrastructureScriptHeader {
	return &DotNetEnvironmentInfrastructureScriptHeader{}
}

func (dnsg DotNetScriptGenerator) GenerateBuildScripts(
	dnd projectmodel.DotNetDeliverable,
	details *jobmodel.JobDetails) []genmodel.ScriptKeyValuePair {
	var results []genmodel.ScriptKeyValuePair
	flattenedDeliverables := dnd.GetFlattenedDeliverables()
	for _, flattenedDeliverable := range *flattenedDeliverables {
		test := dnsg.InfrastructureExists(flattenedDeliverable)
		if !test {
			continue
		}
		for _, buildScript := range dnsg.BuildScripts {
			scriptHeader := NewDotNetBuildScriptHeader(flattenedDeliverable, buildScript)
			details.IncrementTotalProgress()
			templateResult := buildScript.GenerateScriptFromTemplate(scriptHeader)
			result := genmodel.ScriptKeyValuePair{
				KeyElements: flattenedDeliverable.GetScriptKey(buildScript),
				Value:       *templateResult,
				Type:        string(buildScript.Type),
				Extension:   buildScript.Extension,
				ToolScope:   buildScript.ToolScope,
			}
			logging.LogInfoMultiline("Generated build script",
				"Script: "+buildScript.Name,
				"Script Key: "+flattenedDeliverable.GetScriptKeyString(buildScript))
			results = append(results, result)
			details.IncrementProgress()
		}
	}
	return results
}

func (dnsg DotNetScriptGenerator) GenerateBuildInfrastructureScripts(
	dnd projectmodel.DotNetDeliverable,
	details *jobmodel.JobDetails) []genmodel.ScriptKeyValuePair {
	var results []genmodel.ScriptKeyValuePair
	flattenedDeliverables := dnd.GetFlattenedDeliverables()
	for _, flattenedDeliverable := range *flattenedDeliverables {
		test := dnsg.InfrastructureExists(flattenedDeliverable)
		if !test {
			continue
		}
		buildInfrastructure, err := dnsg.persistenceClient.GetBuildInfrastructure(
			flattenedDeliverable.GetRepositoryKey())
		if err != nil {
			panic(err)
		}
		for _, deployScript := range dnsg.DeployScripts {
			scriptHeader := NewDotNetDeployScriptHeader(
				flattenedDeliverable,
				deployScript,
				buildInfrastructure)
			details.IncrementTotalProgress()
			templateResult := deployScript.GenerateScriptFromTemplate(scriptHeader)
			result := genmodel.ScriptKeyValuePair{
				KeyElements: flattenedDeliverable.GetScriptKey(deployScript),
				Value:       *templateResult,
				Type:        string(deployScript.Type),
				Extension:   deployScript.Extension,
				ToolScope:   deployScript.ToolScope,
			}
			logging.LogInfoMultiline("Generated deploy script",
				"Script: "+deployScript.Name,
				"Script Key: "+flattenedDeliverable.GetScriptKeyString(deployScript))
			results = append(results, result)
			details.IncrementProgress()
		}
	}
	return results
}

func (dnsg DotNetScriptGenerator) GenerateBuildDeployScripts(
	dnd projectmodel.DotNetDeliverable,
	details *jobmodel.JobDetails) []genmodel.ScriptKeyValuePair {
	var results []genmodel.ScriptKeyValuePair
	flattenedDeliverables := dnd.GetFlattenedDeliverables()
	for _, flattenedDeliverable := range *flattenedDeliverables {
		test := dnsg.InfrastructureExists(flattenedDeliverable)
		if !test {
			continue
		}
		buildInfrastructure, err := dnsg.persistenceClient.GetBuildInfrastructure(
			flattenedDeliverable.GetRepositoryKey())
		if err != nil {
			panic(err)
		}
		for _, buildDeployScript := range dnsg.BuildDeployScripts {
			scriptHeader := NewDotNetBuildDeployScriptHeader(
				flattenedDeliverable,
				buildDeployScript,
				buildInfrastructure)
			details.IncrementTotalProgress()
			templateResult := buildDeployScript.GenerateScriptFromTemplate(scriptHeader)
			result := genmodel.ScriptKeyValuePair{
				KeyElements: flattenedDeliverable.GetScriptKey(buildDeployScript),
				Value:       *templateResult,
				Type:        string(buildDeployScript.Type),
				Extension:   buildDeployScript.Extension,
				ToolScope:   buildDeployScript.ToolScope,
			}
			logging.LogInfoMultiline("Generated build deploy script",
				"Script: "+buildDeployScript.Name,
				"Script Key: "+flattenedDeliverable.GetScriptKeyString(buildDeployScript))
			results = append(results, result)
			details.IncrementProgress()
		}
	}
	return results
}

func (dnsg DotNetScriptGenerator) InfrastructureExists(
	flattenedDeliverable projectmodel.FlattenedDotNetDeliverable) bool {
	key := flattenedDeliverable.GetRepositoryKey()
	result := linq.From(dnsg.ResourceList.Keys).FirstWithT(func(x inframodel.ResourceKey) bool {
		return key.ProjectName == x.ProjectName &&
			key.SolutionName == x.SolutionName &&
			key.RepositoryName == x.RepositoryName
	})

	if result == nil {
		return false
	}
	return true
}

func (dnsg DotNetScriptGenerator) GenerateEnvironmentInfrastructureScripts() []string {
	var result []string
	scriptHeader := NewDotNetEnvironmentInfrastructureScriptHeader()
	for _, environmentInfraScript := range dnsg.EnvironmentInfrastructureScripts {
		templateResult := environmentInfraScript.GenerateScriptFromTemplate(scriptHeader)
		result = append(result, *templateResult)
	}
	return result
}
