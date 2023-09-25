package postgres

import (
	"context"
	"fmt"
	"time"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/contextutils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AlertHitRepository struct {
	db *gorm.DB
}

func NewAlertHitRepository(db *gorm.DB) *AlertHitRepository {
	return &AlertHitRepository{
		db: db,
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
	if err := a.db.WithContext(ctx).Table("alert_hit").Select("count(*)").Where("time between ? and ?", start, end).Where("config_uuid = ?", id).Take(&result).Error; err != nil {
		contextutils.Logger(ctx).Warnw("cant get rate", "uuid", id, "error", err)
		return 0, fmt.Errorf("get rate: %w", err)
	}

	return result, nil
}

func (a *AlertHitRepository) Create(ctx context.Context, id uuid.UUID) error {
	hit := AlertHit{
		UUID:       uuid.New(),
		ConfigUUID: id,
		Time:       time.Now().UTC(),
	}

	if err := a.db.WithContext(ctx).Table("alert_hit").Create(&hit).Error; err != nil {
		contextutils.Logger(ctx).Warnw("cant create alert hit", "uuid", id, "error", err)
		return fmt.Errorf("create alert hit: %w", err)
	}

	return nil
}

func (a *AlertHitRepository) DeleteAll(ctx context.Context, id uuid.UUID) error {
	if err := a.db.WithContext(ctx).Table("alert_hit").Where("config_uuid = ?", id).Delete(&AlertHit{}).Error; err != nil {
		return fmt.Errorf("delete all hits: %w", err)
	}

	return nil
}
