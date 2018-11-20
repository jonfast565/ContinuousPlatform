package databasemodels

import "github.com/satori/go.uuid"

type IisApplication struct {
	IisApplicationId  uuid.UUID `gorm:"primary_key"`
	ApplicationPoolId uuid.UUID `gorm:"not null"`
	ApplicationName   string    `gorm:"not null"`
	PhysicalPath      string    `gorm:"not null"`
	AuditFields
}
