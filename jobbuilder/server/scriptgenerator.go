package server

import (
	"github.com/jonfast565/continuous-platform/jobbuilder/server/generators"
	"github.com/jonfast565/continuous-platform/models/genmodel"
	"github.com/jonfast565/continuous-platform/models/jobmodel"
	"github.com/jonfast565/continuous-platform/models/projectmodel"
	"github.com/jonfast565/continuous-platform/utilities/logging"
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
		buildDeployScripts := dotNetScriptGenerator.GenerateBuildDeployScripts(*dotNetDeliverable, details)
		scripts = append(scripts, buildScripts...)
		scripts = append(scripts, buildInfraScripts...)
		scripts = append(scripts, buildDeployScripts...)
	}
	return scripts
}
