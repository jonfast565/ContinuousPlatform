package databasemodels

import "time"

type AuditFields struct {
	CreatedDateTime      time.Time `gorm:"not null;default:now()"`
	CreatedBy            string    `gorm:"notnull;default:'SystemCI'"`
	LastModifiedDateTime time.Time `gorm:"not null;default:now()"`
	LastModifiedBy       string    `gorm:"not null;default:'SystemCI'"`
}
