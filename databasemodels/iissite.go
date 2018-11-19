package databasemodels

import (
	"github.com/lib/pq"
	"github.com/satori/go.uuid"
)

type IisSite struct {
	IisSiteId         uuid.UUID
	ApplicationPoolId uuid.UUID
	SiteName          string
	PhysicalPath      string
	SiteApplications  pq.StringArray `gorm:"type:uuid[]"`
	AuditFields
}
