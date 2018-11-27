package server

import (
	"../../logging"
	"../../models/genmodel"
	"../../models/jobmodel"
	"./generators"
)

func GenerateScripts(details *jobmodel.JobDetails) {
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
	scripts := make([]genmodel.ScriptKeyValuePair, 0)
	details.ResetProgress()
	details.SetTotalProgress(int64(len(deliverables.Deliverables)))
	for _, deliverable := range deliverables.Deliverables {
		for _, dotNetDeliverable := range deliverable.DotNetDeliverables {
			buildScripts := dotNetScriptGenerator.GenerateBuildScripts(*dotNetDeliverable)
			buildInfraScripts := dotNetScriptGenerator.GenerateBuildInfrastructureScripts(*dotNetDeliverable)
			scripts = append(scripts, buildScripts...)
			scripts = append(scripts, buildInfraScripts...)
		}
		details.IncrementProgress()
	}

	scriptPackage := genmodel.ScriptPackage{
		Scripts: scripts,
	}

	err = SetScriptCache(scriptPackage)
	if err != nil {
		panic(err)
	}
}
