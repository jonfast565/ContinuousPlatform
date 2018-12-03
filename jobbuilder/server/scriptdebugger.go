package server

import (
	"../../logging"
	"../../models/jobmodel"
	"../../stringutil"
	"io/ioutil"
	"strings"
)

func DeployScriptsForDebugging(details *jobmodel.JobDetails) {
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
		scriptPart := stringutil.ConcatMultipleWithSeparator("-", script.KeyElements...)
		scriptPart = strings.Replace(scriptPart, "/", "-", -1)
		fileName := "C:/Users/***REMOVED***/Desktop/Scripts/" + scriptPart + "-" + script.Type + ".txt"
		logging.LogInfo("Writing " + fileName + " to disk")
		err := ioutil.WriteFile(fileName, []byte(script.Value), 0666)
		if err != nil {
			panic(err)
		}
	}
}
