package genmodel

import (
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
	fileName := debugPathBase + scriptPart + "-" + s.Type + ".txt"
	return fileName
}
