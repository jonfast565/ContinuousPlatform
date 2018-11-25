package genmodel

import (
	"io/ioutil"
)

type ScriptFramework string
type ScriptType string

const (
	DotNet ScriptFramework = "DotNet"
)

const (
	Build                     ScriptType = "Build"
	BuildInfrastructure       ScriptType = "BuildInfrastructure"
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

func (st *ScriptTemplate) LoadTemplateFile() {
	if st.loadedTemplate == "" {
		st.loadedTemplate = st.getTemplateFile()
	}
}

func (st ScriptTemplate) getTemplateFile() string {
	dat, err := ioutil.ReadFile(st.ContentsPath)
	if err != nil {
		panic(err)
	}
	return string(dat)
}
