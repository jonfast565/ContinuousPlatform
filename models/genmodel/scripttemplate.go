package genmodel

import (
	"../../templating"
	"io/ioutil"
)

type ScriptFramework string
type ScriptType string

const (
	DotNet ScriptFramework = "DotNet"
)

const (
	Build                     ScriptType = "Build"
	Deploy                    ScriptType = "Deploy"
	BuildDeploy               ScriptType = "BuildDeploy"
	EnvironmentInfrastructure ScriptType = "EnvironmentInfrastructure"
)

type ScriptTemplate struct {
	Name           string
	Type           ScriptType
	Framework      ScriptFramework
	Extension      string
	ContentsPath   string
	Enabled        bool
	ToolScope      []string
	loadedTemplate string
}

func (st ScriptTemplate) GenerateScriptFromTemplate(scriptHeader interface{}) *string {
	template, err := st.getTemplateFile()
	if err != nil {
		panic(err)
	}
	templateResult, err := templating.RunTemplate(*template, scriptHeader)
	if err != nil {
		panic(err)
	}
	return templateResult
}

func (st ScriptTemplate) getTemplateFile() (*string, error) {
	if st.loadedTemplate == "" {
		dat, err := ioutil.ReadFile(st.ContentsPath)
		if err != nil {
			return nil, err
		}
		st.loadedTemplate = string(dat)
		return &st.loadedTemplate, nil
	} else {
		return &st.loadedTemplate, nil
	}
}
