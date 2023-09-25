package managers

import (
	"context"
	"fmt"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/contextutils"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/ports"
	"github.com/google/uuid"
)

type AlertConfigManager struct {
	alertConfigRepo ports.AlertConfigRepository
}

func NewAlertConfigManager(acr ports.AlertConfigRepository) *AlertConfigManager {
	return &AlertConfigManager{
		alertConfigRepo: acr,
	}
}

func (a *AlertConfigManager) CreateAlertConfig(ctx context.Context, c *models.AlertConfig) (*models.AlertConfig, error) {
	contextutils.Logger(ctx).Infow(
		"creating alert config",
		"message_expression", c.MessageExpression,
		"min_priority", c.MinPriority.String(),
		"duration", c.Duration.String(),
		"number", c.MinRate,
		"comment", c.Comment,
	)

	c, err := a.alertConfigRepo.Create(ctx, c)
	if err != nil {
		contextutils.Logger(ctx).Errorw("cannot create alert config", "error", err)
		return nil, fmt.Errorf("create alert config: %w", err)
	}

	contextutils.Logger(ctx).Infow(
		"created alert config",
		"uuid", c.UUID.String(),
	)

	return c, nil
}

func (a *AlertConfigManager) GetAlertConfigList(ctx context.Context) ([]*models.AlertConfig, error) {
	contextutils.Logger(ctx).Infow("getting alert configs")

	configs, err := a.alertConfigRepo.GetList(ctx)
	if err != nil {
		contextutils.Logger(ctx).Errorw("cannot get alert configs", "error", err)
		return nil, fmt.Errorf("get alert configs: %w", err)
	}

	contextutils.Logger(ctx).Infow("got alert configs")

	return configs, nil
}

func (a *AlertConfigManager) DeleteAlertConfig(ctx context.Context, id uuid.UUID) error {
	contextutils.Logger(ctx).Infow(
		"deleting alert config",
		"uuid", id,
	)

	if err := a.alertConfigRepo.Delete(ctx, id); err != nil {
		contextutils.Logger(ctx).Errorw("cannot delete alert config", "uuid", id, "error", err)
		return fmt.Errorf("delete alert config: %w", err)
	}

	contextutils.Logger(ctx).Infow(
		"deleted alert config",
		"uuid", id,
	)

	return nil
}
