package persistenceclient

import (
	"../../compressutil"
	"../../constants"
	"../../jsonutil"
	"../../models/persistmodel"
	"../../webutil"
	"bytes"
	"encoding/json"
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
	client        http.Client
}

func NewPersistenceClient() PersistenceClient {
	var config ClientConfiguration
	jsonutil.DecodeJsonFromFile(SettingsFilePath, &config)
	return PersistenceClient{configuration: config, client: http.Client{}}
}

func (pc PersistenceClient) GetKeyValueCache(key string) ([]byte, error) {
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
	err = webutil.ExecuteRequestAndReadJsonBody(&pc.client, request, &value)
	if err != nil {
		return nil, err
	}

	uncompressedBytes, err := compressutil.Uncompress(value.Value)
	if err != nil {
		return nil, err
	}

	return uncompressedBytes, nil
}

func (pc PersistenceClient) SetKeyValueCache(key string, value []byte) error {
	// compress payload for speed
	result, err := compressutil.Compress(value)
	if err != nil {
		return err
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

	webutil.ExecuteRequestWithoutRead(&pc.client, request)
	return nil
}
