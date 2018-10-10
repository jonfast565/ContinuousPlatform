package utilities

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Empty struct {
}

func DecodeJsonFromFile(path string, object interface{}) {
	file, _ := os.OpenFile(path, os.O_RDONLY, 0666)
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(object)
	if err != nil {
		panic(err)
	}
	log.Print("Json at " + path + " read successfully.")
}

func DecodeJsonFromBody(r *http.Request, object interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&object)
	if err != nil {
		return err
	}
	return nil
}

func EncodeJsonToBytes(object interface{}) (*[]byte, error) {
	result, err := json.Marshal(&object)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
