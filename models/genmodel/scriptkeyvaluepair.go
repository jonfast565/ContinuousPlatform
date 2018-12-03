package genmodel

import (
	"../../models/jenkinsmodel"
	"../../stringutil"
	"strings"
)

type ScriptKeyValuePair struct {
	KeyElements []string
	Value       string
	Type        string
	Extension   string
	ToolScope   []string
}

func (s ScriptKeyValuePair) GetDebugFilePath(debugPathBase string) string {
	scriptPart := stringutil.ConcatMultipleWithSeparator("-", s.KeyElements...)
	scriptPart = strings.Replace(scriptPart, "/", "-", -1)
	// TODO: Replace with path algos
	fileName := debugPathBase + scriptPart + "-" + s.Type + "." + s.Extension
	return fileName
}

func (s ScriptKeyValuePair) GetJenkinsKeySet() []jenkinsmodel.JenkinsJobKey {
	scriptMetadataKeys := make([]jenkinsmodel.JenkinsJobKey, 0)
	for i := range s.KeyElements {
		if i != len(s.KeyElements)-1 {
			scriptMetadataKeys = append(scriptMetadataKeys, jenkinsmodel.JenkinsJobKey{
				Keys: s.KeyElements[0 : i+1],
				Type: jenkinsmodel.Folder,
			})
		} else {
			scriptMetadataKeys = append(scriptMetadataKeys, jenkinsmodel.JenkinsJobKey{
				Keys: s.KeyElements[0 : i+1],
				Type: jenkinsmodel.PipelineJob,
			})
		}
	}
	return scriptMetadataKeys
}
