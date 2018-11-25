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

func NewDotNetBuildScriptHeader(dnd projectmodel.DotNetDeliverable) *DotNetBuildScriptHeader {
	return &DotNetBuildScriptHeader{}
}

type DotNetBuildInfrastructureScriptHeader struct {
}

func NewDotNetBuildInfrastructureScriptHeader(dnd projectmodel.DotNetDeliverable) *DotNetBuildInfrastructureScriptHeader {
	return &DotNetBuildInfrastructureScriptHeader{}
}

type DotNetEnvironmentInfrastructureScriptHeader struct {
}

func NewDotNetEnvironmentInfrastructureScriptHeader(dnd projectmodel.DotNetDeliverable) *DotNetEnvironmentInfrastructureScriptHeader {
	return &DotNetEnvironmentInfrastructureScriptHeader{}
}

func (dnsg DotNetScriptGenerator) GenerateBuildScripts(dnd projectmodel.DotNetDeliverable) []string {
	var result []string
	scriptHeader := NewDotNetBuildScriptHeader(dnd)
	for _, buildScript := range dnsg.BuildScripts {
		templateResult := buildScript.GenerateScriptFromTemplate(scriptHeader)
		result = append(result, *templateResult)
	}
	return result
}

func (dnsg DotNetScriptGenerator) GenerateBuildInfrastructureScripts(dnd projectmodel.DotNetDeliverable) []string {
	var result []string
	scriptHeader := NewDotNetBuildInfrastructureScriptHeader(dnd)
	for _, buildInfraScript := range dnsg.BuildInfrastructureScripts {
		templateResult := buildInfraScript.GenerateScriptFromTemplate(scriptHeader)
		result = append(result, *templateResult)
	}
	return result
}

func (dnsg DotNetScriptGenerator) GenerateEnvironmentInfrastructureScripts(dnd projectmodel.DotNetDeliverable) []string {
	var result []string
	scriptHeader := NewDotNetEnvironmentInfrastructureScriptHeader(dnd)
	for _, environmentInfraScript := range dnsg.EnvironmentInfrastructureScripts {
		templateResult := environmentInfraScript.GenerateScriptFromTemplate(scriptHeader)
		result = append(result, *templateResult)
	}
	return result
}
