package jsonutil

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Empty struct {
}

// Decodes a json object from a file, used for configuration
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

// Decodes json from a request body, used a lot
func DecodeJsonFromBody(r *http.Request, object interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&object)
	if err != nil {
		return err
	}
	return nil
}

// Turns a JSON value from any object into a byte array
// useful for converting to an ArrayBuffer before the frontend,
// or could be used with ProtoBuf
func EncodeJsonToBytes(object interface{}) (*[]byte, error) {
	result, err := json.Marshal(&object)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Pretty prints json, used in debugging
// TODO: Use?
func JsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}
