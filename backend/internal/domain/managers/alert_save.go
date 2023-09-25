package managers

import (
	"context"
	"fmt"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/contextutils"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/ports"
)

type AlertSaveManager struct {
	repo ports.AlertCreationRepository
}

func NewAlertSaveManager(r ports.AlertCreationRepository) *AlertSaveManager {
	return &AlertSaveManager{
		repo: r,
	}
}

func (a *AlertSaveManager) Save(ctx context.Context, alert *models.Alert) error {
	contextutils.Logger(ctx).Infow(
		"saving alert",
		"uuid", alert.UUID.String(),
	)

	if err := a.repo.Create(ctx, alert); err != nil {
		contextutils.Logger(ctx).Warnw(
			"cannot mark alert as solved",
			"error", err,
			"uuid", alert.UUID.String(),
		)

		return fmt.Errorf("mark solved: %w", err)
	}

	return nil
}
