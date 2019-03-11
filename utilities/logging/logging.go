// A logging package that wraps logging messages for servers/clients
package logging

import (
	"encoding/json"
	"fmt"
	"github.com/go-errors/errors"
	"github.com/jonfast565/continuous-platform/constants"
	"github.com/jonfast565/continuous-platform/utilities/timeutil"
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

// Create a logger
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
	// TODO: Options are only useful for debugging, maybe use a DEBUG flag?
	// log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// Logs the application header to the log file
// Useful for identifying when the application started
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

// Logs application start time
func LogApplicationStart() {
	log.Print("Application started @ " + timeutil.GetCurrentTime() + "...")
}

// Logs the application end time
func LogApplicationEnd() {
	log.Print("Application finished @ " + timeutil.GetCurrentTime() + ".")
}

// Logs that the service has started serving content
func LogContentService(portString string) {
	log.Print("Serving content @ " + portString)
}

// Logs an info message
func LogInfo(logLine string) {
	log.Printf("[Info] %s", logLine)
}

// Logs a fatal message and prints a stack trace
func LogFatal(error interface{}) {
	log.Printf("[Jack the Ripper] %s", error)
	log.Printf("[Stack Dump] %s", errors.Wrap(error, 0).ErrorStack())
}

// Logs a panic/recover series and prints a stack trace
func LogPanicRecover(error interface{}) {
	log.Printf("[Error] %s", error)
	log.Printf("[Stack Dump] %s", errors.Wrap(error, 0).ErrorStack())
}

// Logs a soft error, one that can be ignored
func LogSoftError(errorMessage string, err error) {
	log.Printf("[Soft Error] %s %s", errorMessage, err.Error())
}

// Logs info in a multiline fashion, can accept multiple lines at once
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

// Logs an error, with a specific emphasis on json serialization
// as this type of error is hard to debug
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

// Logs a pretty map
// TODO: In use?
func LogMapPretty(message string, theMap map[string]interface{}) {
	b, err := json.MarshalIndent(theMap, "", "  ")
	if err != nil {
		panic(err)
	}
	LogInfo(message)
	lines := strings.Split(string(b), "\n")
	LogInfoMultiline(lines...)
}
