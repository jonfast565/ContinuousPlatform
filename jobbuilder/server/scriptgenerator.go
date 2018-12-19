package server

import (
	"../../logging"
	"../../models/genmodel"
	"../../models/jobmodel"
	"../../models/projectmodel"
	"./generators"
)

func GenerateScripts(details *jobmodel.JobDetails) bool {
	defer func() {
		if r := recover(); r != nil {
			details.SetJobErrored()
			logging.LogPanicRecover(r)
		}
	}()

	deliverables, err := GetDeliverablesCache()
	if err != nil {
		panic(err)
	}

	dotNetScriptGenerator := generators.NewDotNetScriptGenerator()
	details.ResetProgress()

	var scripts []genmodel.ScriptKeyValuePair
	for _, deliverable := range deliverables.Deliverables {
		deliverableScripts := generateDotNetScripts(details, deliverable, dotNetScriptGenerator)
		scripts = append(scripts, deliverableScripts...)
	}

	scriptPackage := genmodel.ScriptPackage{
		Scripts: scripts,
	}

	err = SetScriptCache(scriptPackage)
	if err != nil {
		panic(err)
	}

	return true
}

func generateDotNetScripts(
	details *jobmodel.JobDetails,
	deliverable projectmodel.Deliverable,
	dotNetScriptGenerator *generators.DotNetScriptGenerator) []genmodel.ScriptKeyValuePair {
	var scripts []genmodel.ScriptKeyValuePair
	for _, dotNetDeliverable := range deliverable.DotNetDeliverables {
		buildScripts := dotNetScriptGenerator.GenerateBuildScripts(*dotNetDeliverable, details)
		buildInfraScripts := dotNetScriptGenerator.GenerateBuildInfrastructureScripts(*dotNetDeliverable, details)
		scripts = append(scripts, buildScripts...)
		scripts = append(scripts, buildInfraScripts...)
	}
	return scripts
}
