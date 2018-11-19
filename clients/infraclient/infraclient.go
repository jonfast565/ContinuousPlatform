package infraclient

import (
	"../../constants"
	"../../jsonutil"
	"../../models/inframodel"
	"../../webutil"
	"net/http"
	"strconv"
)

var (
	SettingsFilePath = "./infraclient-settings.json"
)

type ClientConfiguration struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
}

type InfraClient struct {
	configuration ClientConfiguration
	client        http.Client
}

func NewInfraClient() InfraClient {
	var config ClientConfiguration
	jsonutil.DecodeJsonFromFile(SettingsFilePath, &config)
	return InfraClient{configuration: config, client: http.Client{Timeout: constants.ClientTimeout}}
}

func (ic InfraClient) GetInfrastructure() (*inframodel.InfrastructureMetadata, error) {
	// build service url
	myUrl := webutil.NewEmptyUrl()
	myUrl.SetBase(constants.DefaultScheme,
		ic.configuration.Hostname,
		strconv.Itoa(ic.configuration.Port))
	myUrl.AppendPathFragments([]string{"Daemon", "GetInfrastructureMetadata"})

	// execute request
	request, err := http.NewRequest(constants.PostMethod,
		myUrl.GetUrlStringValue(),
		nil)
	if err != nil {
		return nil, err
	}

	var value inframodel.InfrastructureMetadata
	err = webutil.ExecuteRequestAndReadJsonBody(&ic.client, request, &value)
	if err != nil {
		return nil, err
	}

	return &value, nil
}
