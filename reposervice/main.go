package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var config *Configuration
var postMethod = "POST"

type Configuration struct {
	Port int `json:"port"`
}

func getCurrentTime() string {
	return time.Now().Format(time.RFC850)
}

func getLocalPort(config *Configuration) string {

}

func readConfig(path string) *Configuration {
	file, _ := os.Open(path)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := new(Configuration)
	err := decoder.Decode(&configuration)
	if err != nil {
		panic(err)
	}
	log.Print("Configuration read successfully.")
	return configuration
}

func logError(err error) {
	if terr, ok := err.(*json.UnmarshalTypeError); ok {
		log.Printf("Failed to unmarshal field %s \n", terr.Field)
	} else if terr, ok := err.(*json.InvalidUnmarshalError); ok {
		log.Printf("Failed to unmarshal object %s \n", terr.Error())
	} else {
		log.Println(err)
	}
}

func main() {
	log.Printf("Application started at: %s\n", getCurrentTime())

	file, err := os.OpenFile("error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)
	log.Printf("Log file created.")
	defer file.Close()

	config = readConfig("./appsettings.json")
	router := mux.NewRouter()
	router.HandleFunc("/repositories", getRepositories).Methods(postMethod)
	port := getLocalPort(config)

	log.Print("Serving content...")
	log.Fatal(http.ListenAndServe(port, router))
	log.Printf("Application stopped at: %s\n", getCurrentTime())
}

func getRepositories(w http.ResponseWriter, r *http.Request) {

}
