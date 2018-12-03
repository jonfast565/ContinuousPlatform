package generators

import (
	"../../../clients/persistenceclient"
	"../../../logging"
	"../../../models/genmodel"
	"../../../models/inframodel"
	"../../../models/jobmodel"
	"../../../models/projectmodel"
	"github.com/ahmetb/go-linq"
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
}

func NewDotNetBuildScriptHeader(dnd projectmodel.FlattenedDotNetDeliverable) *DotNetBuildScriptHeader {
	return &DotNetBuildScriptHeader{}
}

type DotNetBuildInfrastructureScriptHeader struct {
}

func NewDotNetBuildInfrastructureScriptHeader(dnd projectmodel.FlattenedDotNetDeliverable) *DotNetBuildInfrastructureScriptHeader {
	return &DotNetBuildInfrastructureScriptHeader{}
}

type DotNetEnvironmentInfrastructureScriptHeader struct {
}

func NewDotNetEnvironmentInfrastructureScriptHeader(
/*env inframodel.EnvironmentPart*/ ) *DotNetEnvironmentInfrastructureScriptHeader {
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
		scriptHeader := NewDotNetBuildScriptHeader(flattenedDeliverable)
		for _, buildScript := range dnsg.BuildScripts {
			details.IncrementTotalProgress()
			templateResult := buildScript.GenerateScriptFromTemplate(scriptHeader)
			result := genmodel.ScriptKeyValuePair{
				KeyElements: flattenedDeliverable.GetScriptKey(),
				Value:       *templateResult,
				Type:        string(buildScript.Type),
				Extension:   buildScript.Extension,
				ToolScope:   buildScript.ToolScope,
			}
			logging.LogInfoMultiline("Generating build script",
				"Script: "+buildScript.Name,
				"Script Key: "+flattenedDeliverable.GetScriptKeyString())
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
		scriptHeader := NewDotNetBuildInfrastructureScriptHeader(flattenedDeliverable)
		for _, buildInfraScript := range dnsg.BuildInfrastructureScripts {
			details.IncrementTotalProgress()
			templateResult := buildInfraScript.GenerateScriptFromTemplate(scriptHeader)
			result := genmodel.ScriptKeyValuePair{
				KeyElements: flattenedDeliverable.GetScriptKey(),
				Value:       *templateResult,
				Type:        string(buildInfraScript.Type),
				Extension:   buildInfraScript.Extension,
				ToolScope:   buildInfraScript.ToolScope,
			}
			logging.LogInfoMultiline("Generated build infrastructure script",
				"Script: "+buildInfraScript.Name,
				"Script Key: "+flattenedDeliverable.GetScriptKeyString())
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
