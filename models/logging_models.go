package models

type LogRecord struct {
	ApplicationName string
	LogLevel        string
	Message         string
	MachineName     string
}

// obsolete?
func NewLogRecordFromParameters(machineName string,
	applicationName string,
	logLevel string,
	message string) LogRecord {
	return LogRecord{
		ApplicationName: applicationName,
		MachineName:     machineName,
		LogLevel:        logLevel,
		Message:         message,
	}
}
