package utilities

import (
	"encoding/json"
	"log"
)

func LogError(err error) {
	if terr, ok := err.(*json.UnmarshalTypeError); ok {
		log.Printf("Failed to unmarshal field %s \n", terr.Field)
	} else if terr, ok := err.(*json.InvalidUnmarshalError); ok {
		log.Printf("Failed to unmarshal object %s \n", terr.Error())
	} else {
		log.Println(err)
	}
}
