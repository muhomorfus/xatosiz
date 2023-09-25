package postgres

import (
	"context"
	"fmt"
	"time"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/contextutils"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AlertConfigRepository struct {
	db *gorm.DB
}

func NewAlertConfigRepository(db *gorm.DB) *AlertConfigRepository {
	return &AlertConfigRepository{
		db: db,
	}
}

func (a *AlertConfigRepository) Create(ctx context.Context, c *models.AlertConfig) (*models.AlertConfig, error) {
	cfg := AlertConfig{
		UUID:              uuid.New(),
		MessageExpression: c.MessageExpression,
		MinPriority:       c.MinPriority.String(),
		Duration:          c.Duration.String(),
		MinRate:           c.MinRate,
		Comment:           c.Comment,
	}

	if err := a.db.WithContext(ctx).Table("alert_config").Create(&cfg).Error; err != nil {
		contextutils.Logger(ctx).Warnw("cant create alert config", "error", err)
		return nil, fmt.Errorf("create alert config: %w", err)
	}

	return &models.AlertConfig{
		UUID:              cfg.UUID,
		MessageExpression: c.MessageExpression,
		MinPriority:       c.MinPriority,
		Duration:          c.Duration,
		MinRate:           c.MinRate,
		Comment:           c.Comment,
	}, nil
}

func (a *AlertConfigRepository) GetList(ctx context.Context) ([]*models.AlertConfig, error) {
	var list []AlertConfig

	if err := a.db.WithContext(ctx).Table("alert_config").Find(&list).Error; err != nil {
		contextutils.Logger(ctx).Warnw("cant get alert configs", "error", err)
		return nil, fmt.Errorf("get alert configs: %w", err)
	}

	result := make([]*models.AlertConfig, len(list))
	for i, e := range list {
		duration, err := time.ParseDuration(e.Duration)
		if err != nil {
			contextutils.Logger(ctx).Warnw("cant convert duration", "error", err, "duration", e.Duration)
			return nil, fmt.Errorf("parse duration: %w", err)
		}

		result[i] = &models.AlertConfig{
			UUID:              e.UUID,
			MessageExpression: e.MessageExpression,
			MinPriority:       models.PriorityFromString(e.MinPriority),
			Duration:          duration,
			MinRate:           e.MinRate,
			Comment:           e.Comment,
		}
	}

	return result, nil
}

func (a *AlertConfigRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := a.db.WithContext(ctx).Table("alert_config").Where("uuid = ?", id).Delete(&AlertConfig{}).Error; err != nil {
		contextutils.Logger(ctx).Warnw("cant delete alert config", "error", err, "uuid", id)
		return fmt.Errorf("delete alert config: %w", err)
	}

	return nil
}
