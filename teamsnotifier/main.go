package main

import (
	"../constants"
	"../jsonutil"
	"../logging"
	"../networking"
	"./server"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var config server.Configuration

func main() {
	logging.CreateLog()
	logging.LogHeader("Teams Notifier")
	logging.LogApplicationStart()

	jsonutil.DecodeJsonFromFile("./appsettings.json", &config)
	router := mux.NewRouter()
	router.HandleFunc("/message", sendMessage).Methods(constants.PostMethod)

	localPort := networking.GetLocalPort(config.Port)
	logging.LogContentService(localPort)
	logging.LogFatal(http.ListenAndServe(localPort, router))
	logging.LogApplicationEnd()
}

func sendMessage(w http.ResponseWriter, r *http.Request) {
	var inputMessage server.InputMessage
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&inputMessage)
	if err != nil {
		logging.LogError(err)
		w.WriteHeader(500)
		return
	}

	responseBytes, err := server.SendMessage(&inputMessage, &config)
	if err != nil {
		logging.LogError(err)
		w.WriteHeader(500)
		return
	}

	result := string(*responseBytes)
	fmt.Fprintf(w, result)
}
