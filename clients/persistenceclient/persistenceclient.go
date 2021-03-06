package persistenceclient

import (
	"bytes"
	"encoding/json"
	"github.com/jonfast565/continuous-platform/constants"
	"github.com/jonfast565/continuous-platform/models/inframodel"
	"github.com/jonfast565/continuous-platform/models/persistmodel"
	"github.com/jonfast565/continuous-platform/utilities/compressutil"
	"github.com/jonfast565/continuous-platform/utilities/jsonutil"
	"github.com/jonfast565/continuous-platform/utilities/webutil"
	"net/http"
	"strconv"
)

var (
	SettingsFilePath = "./persistenceclient-settings.json"
)

type ClientConfiguration struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
}

type PersistenceClient struct {
	configuration ClientConfiguration
}

func NewPersistenceClient() PersistenceClient {
	var config ClientConfiguration
	jsonutil.DecodeJsonFromFile(SettingsFilePath, &config)
	return PersistenceClient{configuration: config}
}

func (pc PersistenceClient) GetKeyValueCache(key string, uncompress bool) ([]byte, error) {
	// build service url
	myUrl := webutil.NewEmptyUrl()
	myUrl.SetBase(constants.DefaultScheme,
		pc.configuration.Hostname,
		strconv.Itoa(pc.configuration.Port))
	myUrl.AppendPathFragments([]string{"Daemon", "GetKeyValueCache"})

	// execute request
	requestBody := persistmodel.KeyRequest{Key: key}
	requestJson, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	urlString := myUrl.GetUrlStringValue()
	request, err := http.NewRequest(constants.PostMethod,
		urlString,
		bytes.NewReader(requestJson))
	if err != nil {
		return nil, err
	}

	var value persistmodel.KeyValueResult
	err = webutil.ExecuteRequestAndReadJsonBody(request, &value)
	if err != nil {
		return nil, err
	}

	var result []byte

	if uncompress {
		result, err = compressutil.Uncompress(value.Value)
		if err != nil {
			return nil, err
		}
	} else {
		result = value.Value
	}

	return result, nil
}

func (pc PersistenceClient) SetKeyValueCache(key string, value []byte, compress bool) error {
	// compress payload for speed
	var result []byte
	var err error
	if compress {
		result, err = compressutil.Compress(value)
		if err != nil {
			return err
		}
	} else {
		result = value
	}

	// build service url
	myUrl := webutil.NewEmptyUrl()
	myUrl.SetBase(constants.DefaultScheme,
		pc.configuration.Hostname,
		strconv.Itoa(pc.configuration.Port))
	myUrl.AppendPathFragments([]string{"Daemon", "SetKeyValueCache"})

	// execute request
	requestBody := persistmodel.KeyValueRequest{Key: key, Value: result}
	requestJson, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(constants.PostMethod,
		myUrl.GetUrlStringValue(),
		bytes.NewReader(requestJson))
	if err != nil {
		return err
	}

	_ = webutil.ExecuteRequestWithoutRead(request)
	return nil
}

func (pc PersistenceClient) GetBuildInfrastructure(
	key inframodel.ResourceKey) (*inframodel.BuildInfrastructureMetadata, error) {
	// build service url
	myUrl := webutil.NewEmptyUrl()
	myUrl.SetBase(constants.DefaultScheme,
		pc.configuration.Hostname,
		strconv.Itoa(pc.configuration.Port))
	myUrl.AppendPathFragments([]string{"Daemon", "GetBuildInfrastructure"})

	// execute request
	requestBody := key
	requestJson, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	urlString := myUrl.GetUrlStringValue()
	request, err := http.NewRequest(constants.PostMethod,
		urlString,
		bytes.NewReader(requestJson))
	if err != nil {
		return nil, err
	}

	var value inframodel.BuildInfrastructureMetadata
	err = webutil.ExecuteRequestAndReadJsonBody(request, &value)
	if err != nil {
		return nil, err
	}

	return &value, nil
}

func (pc PersistenceClient) GetResourceList() (*inframodel.ResourceList, error) {
	// build service url
	myUrl := webutil.NewEmptyUrl()
	myUrl.SetBase(constants.DefaultScheme,
		pc.configuration.Hostname,
		strconv.Itoa(pc.configuration.Port))
	myUrl.AppendPathFragments([]string{"Daemon", "GetResourceList"})

	// execute request
	urlString := myUrl.GetUrlStringValue()
	request, err := http.NewRequest(constants.PostMethod,
		urlString,
		nil)
	if err != nil {
		return nil, err
	}

	var value inframodel.ResourceList
	err = webutil.ExecuteRequestAndReadJsonBody(request, &value)
	if err != nil {
		return nil, err
	}

	return &value, nil
}
