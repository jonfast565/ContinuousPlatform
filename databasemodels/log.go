package databasemodels

import (
	"github.com/satori/go.uuid"
)

type Log struct {
	LogId           uuid.UUID `gorm:"primary_key"`
	MachineName     string    `gorm:"not null"`
	ApplicationName string    `gorm:"not null;default:'PlatformCI'"`
	LogLevel        string    `gorm:"not null"`
	Message         string    `gorm:"not null"`
	AuditFields
}
