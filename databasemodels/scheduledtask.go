package databasemodels

import "github.com/satori/go.uuid"

type ScheduledTask struct {
	ScheduledTaskId           uuid.UUID
	Name                      string
	BinaryPath                string
	BinaryExecutableName      string
	BinaryExecutableArguments string
	ScheduleType              string
	RepeatInterval            int64
	RepetitionDuration        int64
	ExecutionTimeLimit        int64
	Priority                  int64
	AuditFields
}
