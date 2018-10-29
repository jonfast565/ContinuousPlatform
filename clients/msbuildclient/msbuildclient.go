package msbuildclient

import (
	"net/http"
	"../../jsonutil"
	)

var (
	SettingsFilePath = "./repoclient-settings.json"
)

type ClientConfiguration struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
}

type MsBuildClient struct {
	configuration ClientConfiguration
	client        http.Client
}

func NewMsBuildClient() MsBuildClient {
	var config ClientConfiguration
	jsonutil.DecodeJsonFromFile(SettingsFilePath, &config)
	return MsBuildClient{configuration: config, client: http.Client{}}
}
