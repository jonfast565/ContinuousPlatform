package genmodel

import (
	"github.com/jonfast565/continuous-platform/utilities/jsonutil"
)

const (
	DefaultTemplatePath string = "./scripttemplates.json"
)

type ScriptTemplateList struct {
	Templates []ScriptTemplate
}

func NewScriptTemplateList() ScriptTemplateList {
	var list ScriptTemplateList
	jsonutil.DecodeJsonFromFile(DefaultTemplatePath, &list)
	return list
}
