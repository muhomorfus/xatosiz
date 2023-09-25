package kafka

import (
	"context"
	"fmt"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"github.com/Shopify/sarama"
	"github.com/google/uuid"
)

type AlertCreationRepository struct {
	producer sarama.SyncProducer
}

func NewAlertCreationRepository(p sarama.SyncProducer) *AlertCreationRepository {
	return &AlertCreationRepository{producer: p}
}

func (q *AlertCreationRepository) Create(ctx context.Context, a *models.Alert) error {
	a.UUID = uuid.New()

	_, _, err := q.producer.SendMessage(&sarama.ProducerMessage{
		Topic: "alerts",
		Key:   sarama.StringEncoder(a.UUID.String()),
		Value: &alert{
			UUID:    a.UUID.String(),
			Message: a.Message,
			Time:    a.Time,
			Event: alertEvent{
				UUID:      a.Event.UUID.String(),
				TraceUUID: a.Event.TraceUUID.String(),
				GroupUUID: a.Event.GroupUUID.String(),
				Message:   a.Event.Message,
				Priority:  a.Event.Priority.String(),
				Payload:   a.Event.Payload,
				Fixed:     a.Event.Fixed,
				Time:      a.Event.Time,
			},
		},
	})

	if err != nil {
		return fmt.Errorf("send message: %w", err)
	}

	return nil
}
