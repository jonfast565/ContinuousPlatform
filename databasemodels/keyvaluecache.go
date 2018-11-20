package databasemodels

import (
	"github.com/satori/go.uuid"
)

type AppCache struct {
	AppCacheId  uuid.UUID `gorm:"primary_key"`
	MachineName string    `gorm:"not null"`
	KeyString   string    `gorm:"not null"`
	Value       []byte    `gorm:"not null"`
	ValueType   string    `gorm:"not null"`
	AuditFields
}
