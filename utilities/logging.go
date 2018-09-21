package utilities

import (
	"io"
	"log"
	"os"
)

func createLog() error {
	file, err := os.OpenFile("error.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
		return err
	}

	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)
	log.Printf("Log file created.")
	return nil
}
