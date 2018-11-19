package databasemodels

import "github.com/satori/go.uuid"

type WindowsService struct {
	WindowsServiceId          uuid.UUID
	Name                      string
	BinaryPath                string
	BinaryExecutableName      string
	BinaryExecutableArguments string
	AuditFields
}
