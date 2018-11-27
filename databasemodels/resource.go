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
	Type            string         `gorm:"not null"`
	IisApplications pq.StringArray `gorm:"type:uuid[]"`
	IisSites        pq.StringArray `gorm:"type:uuid[]"`
	ScheduledTasks  pq.StringArray `gorm:"type:uuid[]"`
	WindowsServices pq.StringArray `gorm:"type:uuid[]"`
	AuditFields
}
