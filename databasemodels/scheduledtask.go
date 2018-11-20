package databasemodels

import "github.com/satori/go.uuid"

type ScheduledTask struct {
	ScheduledTaskId           uuid.UUID `gorm:"primary_key"`
	Name                      string    `gorm:"not null"`
	BinaryPath                string    `gorm:"not null"`
	BinaryExecutableName      string    `gorm:"not null"`
	BinaryExecutableArguments string    `gorm:"not null"`
	ScheduleType              string    `gorm:"not null"`
	RepeatInterval            int64     `gorm:"not null"`
	RepetitionDuration        int64     `gorm:"not null"`
	ExecutionTimeLimit        int64     `gorm:"not null"`
	Priority                  int64     `gorm:"not null"`
	AuditFields
}
