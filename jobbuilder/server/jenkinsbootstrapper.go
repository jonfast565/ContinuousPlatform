package server

import (
	"../../clients/jenkinsclient"
	"../../constants"
	"../../logging"
	"../../models/jobmodel"
	"../../stringutil"
)

func DeployJenkinsJobs(details *jobmodel.JobDetails) {
	defer func() {
		if r := recover(); r != nil {
			details.SetJobErrored()
			logging.LogPanicRecover(r)
		}
	}()

	scripts, err := GetScriptCache()
	if err != nil {
		panic(err)
	}

	jenkinsClient := jenkinsclient.NewJenkinsClient()
	metadata, err := jenkinsClient.GetJenkinsMetadata()
	if err != nil {
		panic(err)
	}

	_ = metadata.GetFlattenedKeys()
	for _, script := range scripts.Scripts {
		isJenkinsScript := stringutil.StringArrayContains(script.ToolScope, constants.JenkinsToolName)
		if !isJenkinsScript {
			continue
		}
	}
}
