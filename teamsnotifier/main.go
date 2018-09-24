package main

import (
	"../utilities"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

var config Configuration

func main() {
	utilities.CreateLog()
	utilities.LogHeader("Teams Notifier")
	utilities.LogApplicationStart()

	utilities.DecodeJsonFromFile("./appsettings.json", &config)
	router := mux.NewRouter()
	router.HandleFunc("/message", sendMessage).Methods(utilities.PostMethod)

	localPort := utilities.GetLocalPort(config.Port)
	utilities.LogContentService(localPort)
	log.Fatal(http.ListenAndServe(localPort, router))
	utilities.LogApplicationEnd()
}

func getOutputMessage(message *InputMessage, configuration *Configuration) OutputMessage {
	return OutputMessage{
		RoomId: configuration.RoomId,
		// The newline character is supposed to work, but Cisco is lying about it
		// on their api documentation site. Stick with <br> for now, since it works
		Markdown: strings.Join(message.Message, "<br>"),
	}
}

func sendMessage(w http.ResponseWriter, r *http.Request) {
	var inputMessage InputMessage
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&inputMessage)
	if err != nil {
		utilities.LogError(err)
		w.WriteHeader(500)
		return
	}

	responseBytes, err := SendMessage(&inputMessage, &config)
	if err != nil {
		utilities.LogError(err)
		w.WriteHeader(500)
		return
	}

	result := string(*responseBytes)
	fmt.Fprintf(w, result)
}
