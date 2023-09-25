package consumers

import (
	"fmt"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"github.com/google/uuid"
)

func toAlert(a alert) (*models.Alert, error) {
	id, err := uuid.Parse(a.UUID)
	if err != nil {
		return nil, fmt.Errorf("parse alert uuid: %w", err)
	}

	eventID, err := uuid.Parse(a.Event.UUID)
	if err != nil {
		return nil, fmt.Errorf("parse event uuid: %w", err)
	}

	traceID, err := uuid.Parse(a.Event.TraceUUID)
	if err != nil {
		return nil, fmt.Errorf("parse trace uuid: %w", err)
	}

	groupID, err := uuid.Parse(a.Event.GroupUUID)
	if err != nil {
		return nil, fmt.Errorf("parse group uuid: %w", err)
	}

	return &models.Alert{
		UUID:    id,
		Message: a.Message,
		Time:    a.Time,
		Event: &models.Event{
			UUID:      eventID,
			TraceUUID: traceID,
			GroupUUID: groupID,
			Message:   a.Event.Message,
			Priority:  models.PriorityFromString(a.Event.Priority),
			Payload:   a.Event.Payload,
			Fixed:     a.Event.Fixed,
			Time:      a.Event.Time,
		},
	}, nil
}

func toEvent(e event) (*models.Event, error) {
	eventID, err := uuid.Parse(e.UUID)
	if err != nil {
		return nil, fmt.Errorf("parse event uuid: %w", err)
	}

	traceID, err := uuid.Parse(e.TraceUUID)
	if err != nil {
		return nil, fmt.Errorf("parse trace uuid: %w", err)
	}

	groupID, err := uuid.Parse(e.GroupUUID)
	if err != nil {
		return nil, fmt.Errorf("parse group uuid: %w", err)
	}

	return &models.Event{
		UUID:      eventID,
		TraceUUID: traceID,
		GroupUUID: groupID,
		Message:   e.Message,
		Priority:  models.PriorityFromString(e.Priority),
		Payload:   e.Payload,
		Fixed:     e.Fixed,
		Time:      e.Time,
	}, nil
}
