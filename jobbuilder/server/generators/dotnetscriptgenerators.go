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
		template.LoadTemplateFile()
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

type DotNetBuildInfrastructureScriptHeader struct {
}

func (dnsg DotNetScriptGenerator) GenerateBuildScripts(dnd projectmodel.DotNetDeliverable) []string {
	var result []string
	for _, buildScript := range dnsg.BuildScripts {

	}
	return result
}

func (dnsg DotNetScriptGenerator) GenerateBuildInfrastructureScripts(dnd projectmodel.DotNetDeliverable) []string {
	var result []string
	for _, buildInfraScript := range dnsg.BuildInfrastructureScripts {

	}
	return result
}

func (dnsg DotNetScriptGenerator) GenerateEnvironmentInfrastructureScripts(dnd projectmodel.DotNetDeliverable) []string {
	var result []string
	for _, environmentInfraScript := range dnsg.EnvironmentInfrastructureScripts {

	}
	return result
}
