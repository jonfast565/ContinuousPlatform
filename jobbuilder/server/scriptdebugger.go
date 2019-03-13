package server

import (
	"github.com/jonfast565/continuous-platform/models/jobmodel"
	"github.com/jonfast565/continuous-platform/utilities/fileutil"
	"github.com/jonfast565/continuous-platform/utilities/logging"
)

// TODO: Enable when debugging
var DEBUG_SCRIPTS = true

// Allows scripts to be deployed with extensions to the file system
// This is useful for debugging them before they go to Jenkins
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
			fileName := script.GetDebugFilePath("C:\\Users\\jfast\\Desktop\\Files\\Scripts")
			logging.LogInfo("Writing " + fileName + " to disk")
			fileutil.WriteFile(fileName, script.Value)
		}
	}

	return true
}
