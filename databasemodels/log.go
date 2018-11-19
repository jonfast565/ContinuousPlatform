package databasemodels

import (
	"github.com/satori/go.uuid"
)

type Log struct {
	LogId           uuid.UUID `gorm:"primary_key"`
	MachineName     string
	ApplicationName string `gorm:"default:'PlatformCI'"`
	LogLevel        string
	Message         string
	AuditFields
}
