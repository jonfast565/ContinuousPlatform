package generators

import (
	"../../../clients/persistenceclient"
	"../../../models/genmodel"
	"../../../models/projectmodel"
)

type DotNetScriptGenerator struct {
	persistenceClient                persistenceclient.PersistenceClient
	BuildScripts                     []genmodel.ScriptTemplate
	BuildInfrastructureScripts       []genmodel.ScriptTemplate
	EnvironmentInfrastructureScripts []genmodel.ScriptTemplate
}

func NewDotNetScriptGenerator() *DotNetScriptGenerator {
	templateList := genmodel.NewScriptTemplateList()

	scriptGenerator := DotNetScriptGenerator{
		persistenceClient:                persistenceclient.NewPersistenceClient(),
		BuildScripts:                     []genmodel.ScriptTemplate{},
		BuildInfrastructureScripts:       []genmodel.ScriptTemplate{},
		EnvironmentInfrastructureScripts: []genmodel.ScriptTemplate{},
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

func (dnsg DotNetScriptGenerator) GenerateBuildScripts(dnd projectmodel.DotNetDeliverable) []genmodel.ScriptKeyValuePair {
	var result []genmodel.ScriptKeyValuePair
	flattenedDeliverables := dnd.GetFlattenedDeliverables()
	for _, flattenedDeliverable := range *flattenedDeliverables {
		scriptHeader := NewDotNetBuildScriptHeader(flattenedDeliverable)
		for _, buildScript := range dnsg.BuildScripts {
			templateResult := buildScript.GenerateScriptFromTemplate(scriptHeader)
			result = append(result, genmodel.ScriptKeyValuePair{
				KeyElements: flattenedDeliverable.GetScriptKey(),
				Value:       *templateResult,
				ToolScope:   buildScript.ToolScope,
			})
		}
	}
	return result
}

func (dnsg DotNetScriptGenerator) GenerateBuildInfrastructureScripts(
	dnd projectmodel.DotNetDeliverable) []genmodel.ScriptKeyValuePair {
	var result []genmodel.ScriptKeyValuePair
	flattenedDeliverables := dnd.GetFlattenedDeliverables()
	for _, flattenedDeliverable := range *flattenedDeliverables {
		_, err := dnsg.persistenceClient.GetBuildInfrastructure(flattenedDeliverable.GetRepositoryKey())
		if err != nil {
			panic(err)
		}
		scriptHeader := NewDotNetBuildInfrastructureScriptHeader(flattenedDeliverable)
		for _, buildInfraScript := range dnsg.BuildInfrastructureScripts {
			templateResult := buildInfraScript.GenerateScriptFromTemplate(scriptHeader)
			result = append(result, genmodel.ScriptKeyValuePair{
				KeyElements: flattenedDeliverable.GetScriptKey(),
				Value:       *templateResult,
				ToolScope:   buildInfraScript.ToolScope,
			})
		}
	}
	return result
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
