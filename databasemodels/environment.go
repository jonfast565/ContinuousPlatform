package databasemodels

import (
	"github.com/lib/pq"
	"github.com/satori/go.uuid"
)

type Environment struct {
	EnvironmentId uuid.UUID      `gorm:"primary_key"`
	Name          string         `gorm:"not null"`
	BusinessLine  string         `gorm:"not null"`
	Servers       pq.StringArray `gorm:"type:uuid[]"`
	AuditFields
}
