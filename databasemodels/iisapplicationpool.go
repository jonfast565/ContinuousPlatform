package databasemodels

import "github.com/satori/go.uuid"

type IisApplicationPool struct {
	IisApplicationPoolId uuid.UUID `gorm:"primary_key"`
	Name                 string    `gorm:"not null"`
	ProcessType          string    `gorm:"not null"`
	FrameworkVersion     string    `gorm:"not null"`
	AuditFields
}
