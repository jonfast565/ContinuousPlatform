package main

import (
	"../utilities"
	"../utilities/web"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var messagesEndpoint = "https://api.ciscospark.com/v1/messages"

type Configuration struct {
	Port           int    `json:"port"`
	TeamsAuthToken string `json:"teamsAuthToken"`
	RoomId         string `json:"roomId"`
}

type InputMessage struct {
	Message []string `json:"message"`
}

type OutputMessage struct {
	RoomId   string `json:"roomId"`
	Markdown string `json:"markdown"`
}

func SendMessage(inputMessage *InputMessage, config *Configuration) (*[]byte, error) {
	client := &http.Client{}
	outputMessage := getOutputMessage(inputMessage, config)
	outputBytes, err := json.Marshal(outputMessage)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(outputBytes)
	req, err := http.NewRequest(utilities.PostMethod, messagesEndpoint, reader)
	if err != nil {
		return nil, err
	}

	web.AddBearerToken(req, config.TeamsAuthToken)
	web.AddJsonHeader(req)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	utilities.LogInfoMultiline("Sent Message: ",
		fmt.Sprintf("Room: %s", outputMessage.RoomId),
		fmt.Sprintf("Message: %s", outputMessage.Markdown))

	return &responseBytes, nil
}
