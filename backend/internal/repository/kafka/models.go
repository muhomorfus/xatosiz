package kafka

import (
	"encoding/json"
	"time"
)

type event struct {
	UUID      string            `json:"uuid"`
	TraceUUID string            `json:"trace_uuid"`
	GroupUUID string            `json:"group_uuid"`
	Message   string            `json:"message"`
	Priority  string            `json:"priority"`
	Payload   map[string]string `json:"payload"`
	Fixed     bool              `json:"fixed"`
	Time      time.Time         `json:"time"`

	encoded []byte
	err     error
}

func (e *event) ensureEncoded() {
	if e.encoded == nil && e.err == nil {
		e.encoded, e.err = json.Marshal(e)
	}
}

func (e *event) Length() int {
	e.ensureEncoded()
	return len(e.encoded)
}

func (e *event) Encode() ([]byte, error) {
	e.ensureEncoded()
	return e.encoded, e.err
}

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

	encoded []byte
	err     error
}

func (a *alert) ensureEncoded() {
	if a.encoded == nil && a.err == nil {
		a.encoded, a.err = json.Marshal(a)
	}
}

func (a *alert) Length() int {
	a.ensureEncoded()
	return len(a.encoded)
}

func (a *alert) Encode() ([]byte, error) {
	a.ensureEncoded()
	return a.encoded, a.err
}
