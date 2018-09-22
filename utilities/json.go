package utilities

import (
	"encoding/json"
	"log"
	"os"
)

func DecodeJsonFromFile(path string, object interface{}) {
	file, _ := os.Open(path)
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&object)
	if err != nil {
		panic(err)
	}
	log.Print("Json at " + path + " read successfully.")
}
