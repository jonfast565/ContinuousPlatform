package databasemodels

import (
	"github.com/satori/go.uuid"
)

type AppCache struct {
	AppCacheId  uuid.UUID `gorm:"primary_key"`
	MachineName string
	KeyString   string
	Value       []byte
	ValueType   string
	AuditFields
}
