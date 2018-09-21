package utilities

import (
	"encoding/json"
	"log"
	"os"
)

func decodeJsonFromFile(path string) *interface{} {
	file, _ := os.Open(path)
	defer file.Close()
	decoder := json.NewDecoder(file)
	object := new(interface{})
	err := decoder.Decode(&object)
	if err != nil {
		panic(err)
	}
	log.Print("Json at " + path + " read successfully.")
	return object
}