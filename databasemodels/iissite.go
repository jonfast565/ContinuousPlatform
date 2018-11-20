package databasemodels

import (
	"github.com/lib/pq"
	"github.com/satori/go.uuid"
)

type IisSite struct {
	IisSiteId         uuid.UUID      `gorm:"primary_key"`
	ApplicationPoolId uuid.UUID      `gorm:"not null"`
	SiteName          string         `gorm:"not null"`
	PhysicalPath      string         `gorm:"not null"`
	SiteApplications  pq.StringArray `gorm:"type:uuid[]"`
	AuditFields
}
