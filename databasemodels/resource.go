package databasemodels

import (
	"github.com/lib/pq"
	"github.com/satori/go.uuid"
)

type Resource struct {
	DeliverableId   uuid.UUID      `gorm:"primary_key"`
	RepositoryName  string         `gorm:"not null"`
	SolutionName    string         `gorm:"not null"`
	ProjectName     string         `gorm:"not null"`
	IisApplications pq.StringArray `gorm:"type:uuid[];not null"`
	IisSites        pq.StringArray `gorm:"type:uuid[];not null"`
	ScheduledTasks  pq.StringArray `gorm:"type:uuid[];not null"`
	WindowsServices pq.StringArray `gorm:"type:uuid[];not null"`
	AuditFields
}
