package databasemodels

import (
	"github.com/satori/go.uuid"
)

type Server struct {
	ServerId uuid.UUID `gorm:"primary_key"`
	Name     string
	Type     string
	AuditFields
}
