package clickhouse

import (
	"context"
	"fmt"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/jsonenc"
	"github.com/Shopify/sarama"
	"time"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/contextutils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AlertHitRepository struct {
	db       *gorm.DB
	producer sarama.SyncProducer
}

func NewAlertHitRepository(db *gorm.DB, producer sarama.SyncProducer) *AlertHitRepository {
	return &AlertHitRepository{
		db:       db,
		producer: producer,
	}
}

func (a *AlertHitRepository) Get(ctx context.Context, id uuid.UUID, duration time.Duration) (int, error) {
	end := time.Now().UTC()
	start := end.Add(-duration)

	contextutils.Logger(ctx).Debugw(
		"get alert hit",
		"uuid", id,
		"start", start,
		"end", end,
	)

	var result int
	if err := a.db.WithContext(ctx).Table("alert_hit").Select("count(*)").Where("toUnixTimestamp(time) between ? and ?", start.Unix(), end.Add(time.Second).Unix()).Where("config_uuid = ?", id).Take(&result).Error; err != nil {
		contextutils.Logger(ctx).Warnw("cant get rate", "uuid", id, "error", err)
		return 0, fmt.Errorf("get rate: %w", err)
	}

	return result, nil
}

func (a *AlertHitRepository) Create(_ context.Context, id uuid.UUID) error {
	hit := AlertHitKafka{
		UUID:       uuid.New(),
		ConfigUUID: id,
		Time:       time.Now().UTC().Format(timeFmt),
	}

	_, _, err := a.producer.SendMessage(&sarama.ProducerMessage{
		Topic: "alert_hits",
		Key:   sarama.StringEncoder(hit.UUID.String()),
		Value: jsonenc.New(hit),
	})

	if err != nil {
		return fmt.Errorf("send message: %w", err)
	}

	return nil
}

func (a *AlertHitRepository) DeleteAll(ctx context.Context, id uuid.UUID) error {
	if err := a.db.WithContext(ctx).Table("alert_hit").Where("config_uuid = ?", id).UpdateColumn("config_uuid", uuid.Nil).Error; err != nil {
		return fmt.Errorf("delete all hits: %w", err)
	}

	return nil
}
