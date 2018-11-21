package databasemodels

import "github.com/satori/go.uuid"

type WindowsService struct {
	WindowsServiceId          uuid.UUID `gorm:"primary_key"`
	Name                      string    `gorm:"not null"`
	BinaryPath                string    `gorm:"not null"`
	BinaryExecutableName      string    `gorm:"not null"`
	BinaryExecutableArguments string    `gorm:"not null"`
	AuditFields
}