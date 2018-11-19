package databasemodels

import "github.com/satori/go.uuid"

type IisApplication struct {
	IisApplicationId  uuid.UUID
	ApplicationPoolId uuid.UUID
	ApplicationName   string
	ApplicationAlias  string
	PhysicalPath      string
	AuditFields
}
