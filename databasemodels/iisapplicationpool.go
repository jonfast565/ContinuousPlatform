package databasemodels

import "github.com/satori/go.uuid"

type IisApplicationPool struct {
	IisApplicationPoolId uuid.UUID
	Name                 string
	ProcessType          string
	FrameworkVersion     string
	AuditFields
}
