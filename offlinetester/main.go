package offlinetester

import (
	"../utilities"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

var configuration Configuration

func main() {
	utilities.CreateLog()
	utilities.LogHeader("Jenkins Service")
	utilities.LogApplicationStart()

	utilities.DecodeJsonFromFile("./appsettings.json", &configuration)

	router := mux.NewRouter()
	router.PathPrefix("/").HandlerFunc(handleAny).Methods(utilities.PostMethod)

	localPort := utilities.GetLocalPort(configuration.Port)
	utilities.LogContentService(localPort)
	log.Fatal(http.ListenAndServe(localPort, router))
	utilities.LogApplicationEnd()
}

func handleAny(w http.ResponseWriter, r *http.Request) {
	split := strings.Split(r.URL.Path, "/")
	joined := strings.Join()
}
