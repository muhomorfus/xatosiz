package postgres

import (
	"context"
	"fmt"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/contextutils"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AlertRepository struct {
	db *gorm.DB
}

func NewAlertRepository(db *gorm.DB) *AlertRepository {
	return &AlertRepository{
		db: db,
	}
}

func (a *AlertRepository) Create(ctx context.Context, alert *models.Alert) error {
	alertDB := Alert{
		UUID:      alert.UUID,
		Message:   alert.Message,
		EventUUID: alert.Event.UUID,
		Time:      alert.Time,
		Solved:    false,
	}

	if err := a.db.WithContext(ctx).Table("alert").Create(&alertDB).Error; err != nil {
		contextutils.Logger(ctx).Warnw("cant create alert", "uuid", alert.UUID, "error", err)
		return fmt.Errorf("create alert: %w", err)
	}

	return nil
}

func (a *AlertRepository) GetList(ctx context.Context, solved bool) ([]models.Alert, error) {
	var list []Alert

	if err := a.db.WithContext(ctx).Table("alert").Where("solved = ?", solved).Find(&list).Error; err != nil {
		contextutils.Logger(ctx).Warnw("cant get alerts", "error", err)
		return nil, fmt.Errorf("get alerts: %w", err)
	}

	result := make([]models.Alert, len(list))
	for i, e := range list {
		var ev Event

		if err := a.db.WithContext(ctx).Table("event").Where("uuid = ?", e.EventUUID).Take(&ev).Error; err != nil {
			return nil, fmt.Errorf("get event by id: %w", err)
		}

		event, err := toEvent(ev)
		if err != nil {
			return nil, fmt.Errorf("convert event: %w", err)
		}

		result[i] = models.Alert{
			UUID:    e.UUID,
			Message: e.Message,
			Time:    e.Time,
			Event:   &event,
		}
	}

	return result, nil
}

func (a *AlertRepository) MarkSolved(ctx context.Context, id uuid.UUID) error {
	if err := a.db.WithContext(ctx).Table("alert").Where("uuid = ?", id).UpdateColumn("solved", true).Error; err != nil {
		contextutils.Logger(ctx).Warnw("cant update alert", "error", err)
		return fmt.Errorf("update alert: %w", err)
	}

	return nil
}
