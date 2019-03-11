package genmodel

import (
	"github.com/jonfast565/continuous-platform/models/jenkinsmodel"
	"github.com/jonfast565/continuous-platform/utilities/stringutil"
	"path/filepath"
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
	scriptNameExtension := scriptPart + "-" + s.Type + "." + s.Extension
	scriptResultName := strings.Replace(scriptNameExtension, "/", "-", -1)
	fileName := filepath.Join(debugPathBase, scriptResultName)
	return fileName
}

func (s ScriptKeyValuePair) GetJenkinsKeyList() jenkinsmodel.JenkinsJobKeyList {
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
