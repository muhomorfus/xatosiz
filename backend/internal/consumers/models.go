package consumers

import "time"

type alertEvent struct {
	UUID      string            `json:"uuid"`
	TraceUUID string            `json:"trace_uuid"`
	GroupUUID string            `json:"group_uuid"`
	Message   string            `json:"message"`
	Priority  string            `json:"priority"`
	Payload   map[string]string `json:"payload"`
	Fixed     bool              `json:"fixed"`
	Time      time.Time         `json:"time"`
}

type alert struct {
	UUID    string     `json:"uuid"`
	Message string     `json:"message"`
	Time    time.Time  `json:"time"`
	Event   alertEvent `json:"event"`
}

type event struct {
	UUID      string            `json:"uuid"`
	TraceUUID string            `json:"trace_uuid"`
	GroupUUID string            `json:"group_uuid"`
	Message   string            `json:"message"`
	Priority  string            `json:"priority"`
	Payload   map[string]string `json:"payload"`
	Fixed     bool              `json:"fixed"`
	Time      time.Time         `json:"time"`
}
