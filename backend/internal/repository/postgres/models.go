package postgres

import (
	"time"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"github.com/google/uuid"
)

type Group struct {
	UUID uuid.UUID `gorm:"primaryKey"`
}

type GroupPreview struct {
	UUID  uuid.UUID `gorm:"primaryKey"`
	Start time.Time
	End   time.Time
	Fixed bool
	Title string
}

type Component struct {
	Name     string
	Quantity int
}

type Trace struct {
	UUID       uuid.UUID `gorm:"primaryKey"`
	GroupUUID  uuid.UUID
	ParentUUID *uuid.UUID
	Title      string
	TimeStart  time.Time
	TimeEnd    *time.Time
	Component  string
}

type Event struct {
	UUID      uuid.UUID `gorm:"primaryKey"`
	TraceUUID uuid.UUID
	GroupUUID uuid.UUID
	Time      time.Time
	Priority  string
	Message   string
	Payload   *string
	Fixed     bool
}

func priorityToDomain(p string) models.Priority {
	m := map[string]models.Priority{
		"info":    models.PriorityInfo,
		"warning": models.PriorityWarning,
		"error":   models.PriorityError,
		"fatal":   models.PriorityFatal,
	}

	return m[p]
}

type AlertHit struct {
	UUID       uuid.UUID `gorm:"primaryKey"`
	ConfigUUID uuid.UUID
	Time       time.Time
}

type AlertConfig struct {
	UUID              uuid.UUID `gorm:"primaryKey"`
	MessageExpression string
	MinPriority       string
	Duration          string
	MinRate           int
	Comment           string
}

type Alert struct {
	UUID      uuid.UUID `gorm:"primaryKey"`
	Message   string
	EventUUID uuid.UUID
	Time      time.Time
	Solved    bool
}
