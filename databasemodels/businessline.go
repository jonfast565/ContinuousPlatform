package databasemodels

import "github.com/satori/go.uuid"

type BusinessLine struct {
	BusinessLineId uuid.UUID `gorm:"primary_key"`
	Name           string
	Description    string
	AuditFields
}
