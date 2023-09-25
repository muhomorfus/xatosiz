package clickhouse

import (
	"context"
	"encoding/json"
	"fmt"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/jsonenc"
	"github.com/Shopify/sarama"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventRepository struct {
	db       *gorm.DB
	producer sarama.SyncProducer
}

func NewEventRepository(db *gorm.DB, producer sarama.SyncProducer) *EventRepository {
	return &EventRepository{db: db, producer: producer}
}

func (e *EventRepository) Create(_ context.Context, event *models.Event) (*models.Event, error) {
	payload, err := json.Marshal(event.Payload)
	if err != nil {
		return nil, fmt.Errorf("marshal: %w", err)
	}
	p := string(payload)

	ev := EventKafka{
		UUID:      uuid.New(),
		TraceUUID: event.TraceUUID,
		GroupUUID: event.GroupUUID,
		Time:      event.Time.Format(timeFmt),
		Priority:  event.Priority.String(),
		Message:   event.Message,
		Payload:   &p,
	}

	_, _, err = e.producer.SendMessage(&sarama.ProducerMessage{
		Topic: "events",
		Key:   sarama.StringEncoder(ev.UUID.String()),
		Value: jsonenc.New(ev),
	})

	if err != nil {
		return nil, fmt.Errorf("send message: %w", err)
	}

	event.UUID = ev.UUID

	return event, nil
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

func (e *EventRepository) Update(_ context.Context, event *models.Event) error {
	payload, err := json.Marshal(event.Payload)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}
	p := string(payload)

	ev := EventKafka{
		UUID:      event.UUID,
		TraceUUID: event.TraceUUID,
		GroupUUID: event.GroupUUID,
		Time:      event.Time.Format(timeFmt),
		Priority:  event.Priority.String(),
		Message:   event.Message,
		Payload:   &p,
		Fixed:     event.Fixed,
	}

	_, _, err = e.producer.SendMessage(&sarama.ProducerMessage{
		Topic: "events",
		Key:   sarama.StringEncoder(ev.UUID.String()),
		Value: jsonenc.New(ev),
	})

	if err != nil {
		return fmt.Errorf("send message: %w", err)
	}

	event.UUID = ev.UUID

	return nil
}
