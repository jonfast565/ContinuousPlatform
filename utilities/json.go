package utilities

import (
	"bufio"
	"encoding/json"
	"io"
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

func EncodeJsonFromBytes(bytes []byte) string {
	buffer := ByteBuffer.wrap(bytes)
	writer := bufio.NewWriter()
	encoder := json.NewEncoder(writer)
}
