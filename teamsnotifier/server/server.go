// Teams Notifier server
package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jonfast565/continuous-platform/constants"
	"github.com/jonfast565/continuous-platform/utilities/logging"
	"github.com/jonfast565/continuous-platform/utilities/webutil"
	"io/ioutil"
	"net/http"
	"strings"
)

var messagesEndpoint = "https://api.ciscospark.com/v1/messages"

// Configuration for this service
type Configuration struct {
	Port           int    `json:"port"`
	TeamsAuthToken string `json:"teamsAuthToken"`
	RoomId         string `json:"roomId"`
}

// Message that goes into Teams
type InputMessage struct {
	Message []string `json:"message"`
}

// Message that has been converted to Markdown and has had the room id appended to it
type OutputMessage struct {
	RoomId   string `json:"roomId"`
	Markdown string `json:"markdown"`
}

// Send a teams message using an input message and a configuration
func SendMessage(inputMessage *InputMessage, config *Configuration) (*[]byte, error) {
	client := webutil.NewHttpClient()
	outputMessage := getOutputMessage(inputMessage, config)
	outputBytes, err := json.Marshal(outputMessage)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(outputBytes)
	req, err := http.NewRequest(constants.PostMethod, messagesEndpoint, reader)
	if err != nil {
		return nil, err
	}

	webutil.AddBearerToken(req, config.TeamsAuthToken)
	webutil.AddJsonHeader(req)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	logging.LogInfoMultiline("Sent Message: ",
		fmt.Sprintf("Room: %s", outputMessage.RoomId),
		fmt.Sprintf("Message: %s", outputMessage.Markdown))

	return &responseBytes, nil
}

func getOutputMessage(message *InputMessage, configuration *Configuration) OutputMessage {
	return OutputMessage{
		RoomId: configuration.RoomId,
		// The newline character is supposed to work, but Cisco is lying about it
		// on their api documentation site. Stick with <br> for now, since it works
		Markdown: strings.Join(message.Message, "<br>"),
	}
}
