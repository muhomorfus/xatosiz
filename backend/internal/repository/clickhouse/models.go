package clickhouse

import (
	"time"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"github.com/google/uuid"
)

const timeFmt = "2006-01-02 15:04:05.000000000"

type Group struct {
	UUID uuid.UUID `gorm:"primaryKey" json:"uuid"`
}

type GroupPreview struct {
	UUID  uuid.UUID `gorm:"primaryKey" json:"uuid"`
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
	ParentUUID uuid.UUID `gorm:"type:Nullable(UUID)"`
	Title      string
	TimeStart  time.Time
	TimeEnd    *time.Time
	Component  string
}

type TraceKafka struct {
	UUID       uuid.UUID `gorm:"primaryKey" json:"uuid"`
	GroupUUID  uuid.UUID `json:"group_uuid"`
	ParentUUID uuid.UUID `json:"parent_uuid"`
	Title      string    `json:"title"`
	TimeStart  string    `json:"time_start"`
	TimeEnd    *string   `json:"time_end"`
	Component  string    `json:"component"`
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

type EventKafka struct {
	UUID      uuid.UUID `gorm:"primaryKey" json:"uuid"`
	TraceUUID uuid.UUID `json:"trace_uuid"`
	GroupUUID uuid.UUID `json:"group_uuid"`
	Time      string    `json:"time"`
	Priority  string    `json:"priority"`
	Message   string    `json:"message"`
	Payload   *string   `json:"payload"`
	Fixed     bool      `json:"fixed"`
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

type AlertHitKafka struct {
	UUID       uuid.UUID `gorm:"primaryKey" json:"uuid"`
	ConfigUUID uuid.UUID `json:"config_uuid"`
	Time       string    `json:"time"`
}

type Alert struct {
	UUID      uuid.UUID `gorm:"primaryKey"`
	Message   string
	EventUUID uuid.UUID
	Time      time.Time
	Solved    bool
}
