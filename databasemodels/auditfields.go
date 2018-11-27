package databasemodels

import (
	"time"
)

type AuditFields struct {
	CreatedDateTime      time.Time `gorm:"not null;default:now()"`
	CreatedBy            string    `gorm:"not null;default:'ContinuousPlatform'"`
	LastModifiedDateTime time.Time `gorm:"not null;default:now()"`
	LastModifiedBy       string    `gorm:"not null;default:'ContinuousPlatform'"`
}
