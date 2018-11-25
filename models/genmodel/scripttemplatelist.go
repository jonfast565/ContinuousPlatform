package genmodel

import (
	"../../jsonutil"
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
