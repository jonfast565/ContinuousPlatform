package databasemodels

import (
	"github.com/lib/pq"
	"github.com/satori/go.uuid"
)

type Resource struct {
	DeliverableId   uuid.UUID `gorm:"primary_key"`
	RepositoryName  string
	SolutionName    string
	ProjectName     string
	IisApplications pq.StringArray `gorm:"type:uuid[]"`
	IisSites        pq.StringArray `gorm:"type:uuid[]"`
	ScheduledTasks  pq.StringArray `gorm:"type:uuid[]"`
	WindowsServices pq.StringArray `gorm:"type:uuid[]"`
	AuditFields
}
