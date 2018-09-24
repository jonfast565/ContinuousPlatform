package main

import (
	"../utilities"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var messagesEndpoint = "https://api.ciscospark.com/v1/messages"
var config *Configuration

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

func main() {
	utilities.CreateLog()
	utilities.LogHeader("Teams Notifier")
	utilities.LogApplicationStart()

	var configuration Configuration
	utilities.DecodeJsonFromFile("./appsettings.json", &configuration)
	router := mux.NewRouter()
	router.HandleFunc("/message", sendMessage).Methods(utilities.PostMethod)
	port := utilities.GetLocalPort(config.Port)

	log.Print("Serving content...")
	log.Fatal(http.ListenAndServe(port, router))
	utilities.LogApplicationEnd()
}

func getOutputMessage(message *InputMessage, configuration *Configuration) OutputMessage {
	return OutputMessage{
		RoomId: configuration.RoomId,
		// TODO: The newline character is supposed to work, but Cisco is lying about it
		// on their api documentation site. Stick with <br> for now, since it works
		Markdown: strings.Join(message.Message, "<br>"),
	}
}

func sendMessage(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}

	var inputMessage InputMessage
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&inputMessage)
	if err != nil {
		utilities.LogError(err)
		w.WriteHeader(500)
		return
	}

	outputMessage := getOutputMessage(&inputMessage, config)
	outputBytes, err := json.Marshal(outputMessage)
	if err != nil {
		utilities.LogError(err)
		w.WriteHeader(500)
		return
	}

	reader := bytes.NewReader(outputBytes)
	req, err := http.NewRequest(utilities.PostMethod, messagesEndpoint, reader)
	if err != nil {
		utilities.LogError(err)
		w.WriteHeader(500)
		return
	}

	utilities.LogInfo(string(outputBytes))
	utilities.AddBearerToken(req, config.TeamsAuthToken)
	utilities.AddJsonHeader(req)
	resp, err := client.Do(req)

	if err != nil {
		utilities.LogError(err)
		w.WriteHeader(500)
		return
	}

	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		utilities.LogError(err)
		w.WriteHeader(500)
		return
	}

	result := string(responseBytes)
	log.Print(result)

	fmt.Fprintf(w, result)
}
