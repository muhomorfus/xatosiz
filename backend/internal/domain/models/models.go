package models

import (
	"time"

	"github.com/google/uuid"
)

type Priority int

const (
	PriorityInfo Priority = iota
	PriorityWarning
	PriorityError
	PriorityFatal
)

func PriorityFromString(s string) Priority {
	switch s {
	case "info":
		return PriorityInfo
	case "warning":
		return PriorityWarning
	case "error":
		return PriorityError
	case "fatal":
		return PriorityFatal
	default:
		return 0
	}
}

func (p Priority) String() string {
	switch p {
	case PriorityInfo:
		return "info"
	case PriorityWarning:
		return "warning"
	case PriorityError:
		return "error"
	case PriorityFatal:
		return "fatal"
	default:
		return "unknown"
	}
}

type Group struct {
	UUID   uuid.UUID
	Traces []Trace
}

func (g *Group) HasActiveEvents() bool {
	return haveActiveEvents(g.Traces)
}

func haveActiveEvents(traces []Trace) bool {
	if len(traces) == 0 {
		return false
	}

	for _, t := range traces {
		for _, e := range t.Events {
			if !e.Fixed {
				return true
			}
		}

		if haveActiveEvents(t.Children) {
			return true
		}
	}

	return false
}

type TraceBaseInfo struct {
	GroupUUID  *uuid.UUID
	ParentUUID *uuid.UUID
	Title      string
	Component  string
}

type Trace struct {
	UUID       uuid.UUID
	GroupUUID  uuid.UUID
	ParentUUID *uuid.UUID
	Start, End time.Time
	Title      string
	Component  string
	Children   []Trace
	Events     []Event
	Finished   bool
}

type EventBaseInfo struct {
	GroupUUID *uuid.UUID
	TraceUUID *uuid.UUID
	Message   string
	Priority  Priority
	Payload   map[string]string
	Component string
}

type Event struct {
	UUID      uuid.UUID
	TraceUUID uuid.UUID
	GroupUUID uuid.UUID
	Message   string
	Priority  Priority
	Payload   map[string]string
	Fixed     bool
	Time      time.Time
}

type Component struct {
	Name     string
	Quantity int
}

type GroupPreview struct {
	UUID            uuid.UUID
	Title           string
	HasActiveEvents bool
	Events          []Event
	Components      []Component
	Start           time.Time
	End             time.Time
	Duration        time.Duration
}

type Filters struct {
	Limit     int
	Component string
}

type AlertConfig struct {
	UUID              uuid.UUID
	MessageExpression string
	MinPriority       Priority
	Duration          time.Duration
	MinRate           int

	Comment string
}

type Alert struct {
	UUID    uuid.UUID
	Message string
	Time    time.Time
	Event   *Event
}
