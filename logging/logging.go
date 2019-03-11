package logging

import (
	"encoding/json"
	"fmt"
	"github.com/go-errors/errors"
	"github.com/jonfast565/continuous-platform/constants"
	"github.com/jonfast565/continuous-platform/timeutil"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
	"strings"
)

const (
	Filename   = "./error.log"
	MaxSize    = 50
	MaxBackups = 1
	MaxAge     = 10
	Compress   = true
)

func CreateLog() {
	logger := &lumberjack.Logger{
		Filename:   Filename,
		MaxSize:    MaxSize,
		MaxBackups: MaxBackups,
		MaxAge:     MaxAge,
		Compress:   Compress,
	}
	mw := io.MultiWriter(os.Stdout, logger)
	log.SetOutput(mw)
	log.Printf("Log file created.")

	// set options
	// log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func LogHeader(applicationName string) {
	log.Print(constants.Header)
	log.Print(constants.Divider)
	log.Print("Application: " + applicationName)
	log.Print(fmt.Sprintf("Version: %s.%s.%s",
		constants.MajorVersion,
		constants.MinorVersion,
		constants.BuildNumber))
	log.Print(constants.Divider)
}

func LogApplicationStart() {
	log.Print("Application started @ " + timeutil.GetCurrentTime() + "...")
}

func LogApplicationEnd() {
	log.Print("Application finished @ " + timeutil.GetCurrentTime() + ".")
}

func LogContentService(portString string) {
	log.Print("Serving content @ " + portString)
}

func LogInfo(logLine string) {
	log.Printf("[Info] %s", logLine)
}

func LogFatal(error interface{}) {
	log.Printf("[Jack the Ripper] %s", error)
	log.Printf("[Stack Dump] %s", errors.Wrap(error, 0).ErrorStack())
}

func LogPanicRecover(error interface{}) {
	log.Printf("[Error] %s", error)
	log.Printf("[Stack Dump] %s", errors.Wrap(error, 0).ErrorStack())
}

func LogSoftError(errorMessage string, err error) {
	log.Printf("[Soft Error] %s %s", errorMessage, err.Error())
}

func LogInfoMultiline(logLines ...string) {
	result := ""
	for i, line := range logLines {
		if i == 0 {
			result += fmt.Sprintf("[Info] %s\n", line)
		} else {
			result += fmt.Sprintf("     - %s\n", line)
		}
	}
	log.Printf(result)
}

func LogError(err error) {
	if terr, ok := err.(*json.UnmarshalTypeError); ok {
		log.Printf("[Error] failed to unmarshal field %s \n", terr.Field)
	} else if terr, ok := err.(*json.InvalidUnmarshalError); ok {
		log.Printf("[Error] failed to unmarshal object %s \n", terr.Error())
	} else {
		log.Printf("[Error] %s", err.Error())
	}
	log.Printf("[Stack Dump] %s", errors.Wrap(err, 0).ErrorStack())
}

func LogMapPretty(message string, theMap map[string]interface{}) {
	b, err := json.MarshalIndent(theMap, "", "  ")
	if err != nil {
		panic(err)
	}
	LogInfo(message)
	lines := strings.Split(string(b), "\n")
	LogInfoMultiline(lines...)
}
