package utilities

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
)

const (
	Filename   = "./error.log"
	MaxSize    = 50
	MaxBackups = 1
	MaxAge     = 10
	Compress   = true
)

const Header = "\n    _____            __                 __________\n" +
	"   / ___/__  _______/ /____  ____ ___  / ____/  _/\n" +
	"  \\__ \\/ / / / ___/ __/ _ \\/ __ `__ \\/ /    / /\n" +
	" ___/ / /_/ (__  ) /_/  __/ / / / / / /____/ /\n" +
	"/____/\\__, /____/\\__/\\___/_/ /_/ /_/\\____/___/\n" +
	"     /____/\n\n"

const Divider = "-----------------------------------------------"

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
}

func LogHeader(applicationName string) {
	log.Print(Header)
	log.Print(Divider)
	log.Print("Application: " + applicationName)
	log.Print(Divider)
}

func LogApplicationStart() {
	log.Print("Application started...")
}

func LogApplicationEnd() {
	log.Print("Application finished.")
}

func LogContentService(portString string) {
	log.Print("Serving content @ " + portString)
}

func LogInfo(logLine string) {
	log.Printf("[Info] %s", logLine)
}

func LogInfoMultiline(logLines ...string) {
	for i, line := range logLines {
		if i == 0 {
			log.Printf("[Info] %s", line)
		} else {
			log.Printf("     - %s", line)
		}
	}
}
