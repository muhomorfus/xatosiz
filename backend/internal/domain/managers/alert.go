package managers

import (
	"context"
	"fmt"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/contextutils"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/ports"
	"github.com/google/uuid"
)

type AlertManager struct {
	alertRepo ports.AlertRepository
}

func NewAlertManager(ar ports.AlertRepository) *AlertManager {
	return &AlertManager{
		alertRepo: ar,
	}
}

func (a *AlertManager) GetAlertList(ctx context.Context, solved bool) ([]models.Alert, error) {
	contextutils.Logger(ctx).Infow(
		"getting alerts",
		"solved", solved,
	)

	alerts, err := a.alertRepo.GetList(ctx, solved)
	if err != nil {
		contextutils.Logger(ctx).Errorw(
			"cannot get alerts",
			"error", err,
			"solved", solved,
		)

		return nil, fmt.Errorf("get alerts: %w", err)
	}

	return alerts, nil
}

func (a *AlertManager) SolveAlert(ctx context.Context, id uuid.UUID) error {
	contextutils.Logger(ctx).Infow(
		"solving alert",
		"uuid", id.String(),
	)

	if err := a.alertRepo.MarkSolved(ctx, id); err != nil {
		contextutils.Logger(ctx).Warnw(
			"cannot mark alert as solved",
			"error", err,
			"uuid", id.String(),
		)

		return fmt.Errorf("mark solved: %w", err)
	}

	return nil
}
