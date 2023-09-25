package postgres

import (
	"context"
	"encoding/json"
	"fmt"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (e *EventRepository) Create(ctx context.Context, event *models.Event) (*models.Event, error) {
	payload, err := json.Marshal(event.Payload)
	if err != nil {
		return nil, fmt.Errorf("marshal: %w", err)
	}
	p := string(payload)

	ev := Event{
		UUID:      uuid.New(),
		TraceUUID: event.TraceUUID,
		GroupUUID: event.GroupUUID,
		Time:      event.Time,
		Priority:  event.Priority.String(),
		Message:   event.Message,
		Payload:   &p,
	}

	if err := e.db.WithContext(ctx).Table("event").Create(&ev).Error; err != nil {
		return nil, fmt.Errorf("create event: %w", err)
	}

	result, err := toEvent(ev)
	if err != nil {
		return nil, fmt.Errorf("convert event: %w", err)
	}

	return &result, nil
}

func (e *EventRepository) Get(ctx context.Context, id uuid.UUID) (*models.Event, error) {
	var ev Event

	if err := e.db.WithContext(ctx).Table("event").Where("uuid = ?", id).Take(&ev).Error; err != nil {
		return nil, fmt.Errorf("get event by id: %w", err)
	}

	result, err := toEvent(ev)
	if err != nil {
		return nil, fmt.Errorf("convert event: %w", err)
	}

	return &result, nil
}

func (e *EventRepository) GetList(ctx context.Context, fixed bool) ([]*models.Event, error) {
	var el []Event

	if err := e.db.WithContext(ctx).Table("event").Where("fixed = ?", fixed).Find(&el).Error; err != nil {
		return nil, fmt.Errorf("get events: %w", err)
	}

	result := make([]*models.Event, len(el))
	for i, ev := range el {
		converted, err := toEvent(ev)
		if err != nil {
			return nil, fmt.Errorf("convert event: %w", err)
		}

		result[i] = &converted
	}

	return result, nil
}

func (e *EventRepository) Update(ctx context.Context, event *models.Event) error {
	payload, err := json.Marshal(event.Payload)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}
	p := string(payload)

	ev := Event{
		UUID:      event.UUID,
		TraceUUID: event.TraceUUID,
		Time:      event.Time,
		Priority:  event.Priority.String(),
		Message:   event.Message,
		Payload:   &p,
		Fixed:     event.Fixed,
	}

	if err := e.db.WithContext(ctx).Table("event").Save(&ev).Error; err != nil {
		return fmt.Errorf("save event: %w", err)
	}

	return nil
}
