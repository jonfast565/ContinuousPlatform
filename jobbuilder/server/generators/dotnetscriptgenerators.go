package generators

import (
	"../../../clients/persistenceclient"
	"../../../logging"
	"../../../models/genmodel"
	"../../../models/inframodel"
	"../../../models/jobmodel"
	"../../../models/projectmodel"
	"../../../timeutil"
	"github.com/ahmetb/go-linq"
	"github.com/satori/go.uuid"
)

type DotNetScriptGenerator struct {
	persistenceClient                persistenceclient.PersistenceClient
	BuildScripts                     []genmodel.ScriptTemplate
	BuildInfrastructureScripts       []genmodel.ScriptTemplate
	EnvironmentInfrastructureScripts []genmodel.ScriptTemplate
	ResourceList                     inframodel.ResourceList
}

func NewDotNetScriptGenerator() *DotNetScriptGenerator {
	templateList := genmodel.NewScriptTemplateList()

	scriptGenerator := DotNetScriptGenerator{
		persistenceClient:                persistenceclient.NewPersistenceClient(),
		BuildScripts:                     []genmodel.ScriptTemplate{},
		BuildInfrastructureScripts:       []genmodel.ScriptTemplate{},
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
		case genmodel.BuildInfrastructure:
			scriptGenerator.BuildInfrastructureScripts = append(scriptGenerator.BuildInfrastructureScripts, template)
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
	PublishProfiles        []projectmodel.MsBuildPublishProfile
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
		Solutions:              dnd.Project.SolutionParents,
		Project:                *dnd.Project,
		ProjectFolder:          dnd.Project.FolderPath,
		PublishProfiles:        []projectmodel.MsBuildPublishProfile{},
		TargetFrameworks:       dnd.Project.TargetFrameworks,
		DefaultNamespace:       dnd.Project.DefaultNamespace,
		SolutionConfigurations: dnd.Solution.Configurations,
		CanonicalId:            dnd.GetScriptKeyString(template),
		DashedCanonicalId:      dnd.GetScriptKeyString(template),
		Hash:                   uid.String(),
		GeneratedDateTime:      timeutil.GetCurrentTime(),
	}
}

type DotNetBuildInfrastructureScriptHeader struct {
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

func NewDotNetBuildInfrastructureScriptHeader(
	dnd projectmodel.FlattenedDotNetDeliverable,
	template genmodel.ScriptTemplate,
	bim *inframodel.BuildInfrastructureMetadata) *DotNetBuildInfrastructureScriptHeader {
	uid, _ := uuid.NewV4()
	return &DotNetBuildInfrastructureScriptHeader{
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
			logging.LogInfoMultiline("Generating build script",
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
		for _, buildInfraScript := range dnsg.BuildInfrastructureScripts {
			scriptHeader := NewDotNetBuildInfrastructureScriptHeader(
				flattenedDeliverable,
				buildInfraScript,
				buildInfrastructure)
			details.IncrementTotalProgress()
			templateResult := buildInfraScript.GenerateScriptFromTemplate(scriptHeader)
			result := genmodel.ScriptKeyValuePair{
				KeyElements: flattenedDeliverable.GetScriptKey(buildInfraScript),
				Value:       *templateResult,
				Type:        string(buildInfraScript.Type),
				Extension:   buildInfraScript.Extension,
				ToolScope:   buildInfraScript.ToolScope,
			}
			logging.LogInfoMultiline("Generated build infrastructure script",
				"Script: "+buildInfraScript.Name,
				"Script Key: "+flattenedDeliverable.GetScriptKeyString(buildInfraScript))
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
