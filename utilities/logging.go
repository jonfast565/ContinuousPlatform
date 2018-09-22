package utilities

import (
	"io"
	"log"
	"os"
)

func CreateLog() (*os.File, error) {
	file, err := os.OpenFile("error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)
	log.Printf("Log file created.")
	return file, nil
}

func LogApplicationStart() {

}

func LogApplicationEnd() {

}
