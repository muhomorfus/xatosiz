package kafka

import (
	"context"
	"fmt"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"github.com/Shopify/sarama"
)

type AlertProcessingQueue struct {
	producer sarama.SyncProducer
}

func NewAlertProcessingQueue(p sarama.SyncProducer) *AlertProcessingQueue {
	return &AlertProcessingQueue{producer: p}
}

func (q *AlertProcessingQueue) Push(ctx context.Context, e *models.Event) error {
	_, _, err := q.producer.SendMessage(&sarama.ProducerMessage{
		Topic: "alert_processing",
		Key:   sarama.StringEncoder(e.UUID.String()),
		Value: &event{
			UUID:      e.UUID.String(),
			TraceUUID: e.TraceUUID.String(),
			GroupUUID: e.GroupUUID.String(),
			Message:   e.Message,
			Priority:  e.Priority.String(),
			Payload:   e.Payload,
			Fixed:     e.Fixed,
			Time:      e.Time,
		},
	})
	if err != nil {
		return fmt.Errorf("send message to queue: %w", err)
	}

	return nil
}
