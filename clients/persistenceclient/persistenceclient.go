package persistenceclient

import (
	"../../constants"
	"../../jsonutil"
	"../../models/persistmodel"
	"../../webutil"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
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

	request, err := http.NewRequest(constants.GetMethod,
		myUrl.GetUrlStringValue(),
		bytes.NewReader(requestJson))
	if err != nil {
		return nil, err
	}

	var value persistmodel.KeyValueResult
	err = webutil.ExecuteRequestAndReadJsonBody(&pc.client, request, value)
	if err != nil {
		return nil, err
	}

	reader, err := gzip.NewReader(bytes.NewReader(value.Value))
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	uncompressedBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return uncompressedBytes, nil
}

func (pc PersistenceClient) SetKeyValueCache(key string, value []byte) error {
	// compress payload for speed
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	w.Write(value)
	defer w.Close()
	result := b.Bytes()

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

	request, err := http.NewRequest(constants.GetMethod,
		myUrl.GetUrlStringValue(),
		bytes.NewReader(requestJson))
	if err != nil {
		return err
	}

	webutil.ExecuteRequestWithoutRead(&pc.client, request)
	return nil
}
