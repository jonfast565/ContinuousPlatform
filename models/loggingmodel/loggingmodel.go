package loggingmodel

type LogRecord struct {
	ApplicationName string
	LogLevel        string
	Message         string
}

// obsolete?
func NewLogRecordFromParameters(machineName string,
	applicationName string,
	logLevel string,
	message string) LogRecord {
	return LogRecord{
		ApplicationName: applicationName,
		LogLevel:        logLevel,
		Message:         message,
	}
}
