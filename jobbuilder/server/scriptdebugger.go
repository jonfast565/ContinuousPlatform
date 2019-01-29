package server

import (
	"../../fileutil"
	"../../logging"
	"../../models/jobmodel"
)

// TODO: Enable when debugging
var DEBUG_SCRIPTS = true

func DeployScriptsForDebugging(details *jobmodel.JobDetails) bool {
	defer func() {
		if r := recover(); r != nil {
			details.SetJobErrored()
			logging.LogPanicRecover(r)
		}
	}()

	if DEBUG_SCRIPTS {
		scripts, err := GetScriptCache()
		if err != nil {
			panic(err)
		}

		for _, script := range scripts.Scripts {
			fileName := script.GetDebugFilePath("C:\\Users\\***REMOVED***\\Desktop\\Files\\Scripts")
			logging.LogInfo("Writing " + fileName + " to disk")
			fileutil.WriteFile(fileName, script.Value)
		}
	}

	return true
}
