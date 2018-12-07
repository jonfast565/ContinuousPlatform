package server

import (
	"../../fileutil"
	"../../logging"
	"../../models/jobmodel"
)

func DeployScriptsForDebugging(debugBasePath string, details *jobmodel.JobDetails) {
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

	for _, script := range scripts.Scripts {
		fileName := script.GetDebugFilePath(debugBasePath)
		logging.LogInfo("Writing " + fileName + " to disk")
		fileutil.WriteFile(fileName, script.Value)
	}
}
