package server

import (
	"github.com/jonfast565/continuous-platform/fileutil"
	"github.com/jonfast565/continuous-platform/logging"
	"github.com/jonfast565/continuous-platform/models/jobmodel"
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
			fileName := script.GetDebugFilePath("C:\\Users\\jfast\\Desktop\\Files\\Scripts")
			logging.LogInfo("Writing " + fileName + " to disk")
			fileutil.WriteFile(fileName, script.Value)
		}
	}

	return true
}
