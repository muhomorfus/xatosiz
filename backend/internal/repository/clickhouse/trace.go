package clickhouse

import (
	"context"
	"fmt"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/jsonenc"
	"github.com/AlekSi/pointer"
	"github.com/Shopify/sarama"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TraceRepository struct {
	db       *gorm.DB
	producer sarama.SyncProducer
}

func NewTraceRepository(db *gorm.DB, producer sarama.SyncProducer) *TraceRepository {
	return &TraceRepository{db: db, producer: producer}
}

func (t *TraceRepository) Create(_ context.Context, trace *models.Trace) (*models.Trace, error) {
	parentUUID := uuid.Nil
	if trace.ParentUUID != nil {
		parentUUID = *trace.ParentUUID
	}

	tr := TraceKafka{
		UUID:       uuid.New(),
		GroupUUID:  trace.GroupUUID,
		ParentUUID: parentUUID,
		Title:      trace.Title,
		TimeStart:  trace.Start.Format(timeFmt),
		Component:  trace.Component,
	}

	_, _, err := t.producer.SendMessage(&sarama.ProducerMessage{
		Topic: "traces",
		Key:   sarama.StringEncoder(tr.UUID.String()),
		Value: jsonenc.New(tr),
	})

	if err != nil {
		return nil, fmt.Errorf("send message: %w", err)
	}

	trace.UUID = tr.UUID

	return trace, nil
}

func (t *TraceRepository) Update(_ context.Context, trace *models.Trace) error {
	parentUUID := uuid.Nil
	if trace.ParentUUID != nil {
		parentUUID = *trace.ParentUUID
	}

	tr := TraceKafka{
		UUID:       trace.UUID,
		GroupUUID:  trace.GroupUUID,
		ParentUUID: parentUUID,
		Title:      trace.Title,
		TimeStart:  trace.Start.Format(timeFmt),
		TimeEnd:    pointer.To(trace.End.Format(timeFmt)),
		Component:  trace.Component,
	}

	_, _, err := t.producer.SendMessage(&sarama.ProducerMessage{
		Topic: "traces",
		Key:   sarama.StringEncoder(tr.UUID.String()),
		Value: jsonenc.New(tr),
	})

	if err != nil {
		return fmt.Errorf("send message: %w", err)
	}

	return nil
}

func (t *TraceRepository) Get(ctx context.Context, id uuid.UUID) (*models.Trace, error) {
	var tr Trace
	if err := t.db.WithContext(ctx).Table("trace").Where("uuid = ?", id).Take(&tr).Error; err != nil {
		return nil, fmt.Errorf("get trace from db: %w", err)
	}

	trace, err := toTrace(ctx, t.db, tr)
	if err != nil {
		return nil, fmt.Errorf("convert trace: %w", err)
	}

	return &trace, nil
}
