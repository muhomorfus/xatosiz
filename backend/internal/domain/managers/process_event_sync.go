package managers

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/contextutils"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/ports"
)

type ProcessEventSyncManager struct {
	alertHitRepo      ports.AlertHitRepository
	alertConfigRepo   ports.AlertConfigRepository
	alertCreationRepo ports.AlertCreationRepository
}

func NewProcessEventSyncManager(ahr ports.AlertHitRepository, acr ports.AlertConfigRepository, acrr ports.AlertCreationRepository) *ProcessEventSyncManager {
	return &ProcessEventSyncManager{
		alertHitRepo:      ahr,
		alertConfigRepo:   acr,
		alertCreationRepo: acrr,
	}
}

func (a *ProcessEventSyncManager) processAlert(ctx context.Context, e *models.Event, c *models.AlertConfig) (bool, error) {
	// Check regexp of message expression.
	match, err := regexp.MatchString(c.MessageExpression, e.Message)
	if err != nil {
		contextutils.Logger(ctx).Errorw(
			"cannot parse regexp",
			"error", err,
			"regexp", c.MessageExpression,
		)
		return false, fmt.Errorf("parse regexp: %w", err)
	}

	if !match {
		contextutils.Logger(ctx).Infow(
			"pattern not match",
			"event", e.UUID.String(),
			"alert_config", c.UUID.String(),
			"pattern", c.MessageExpression,
			"message", e.Message,
		)

		return false, nil
	}

	contextutils.Logger(ctx).Debugw(
		"pattern matched",
		"event", e.UUID.String(),
		"alert_config", c.UUID.String(),
		"pattern", c.MessageExpression,
		"message", e.Message,
	)

	// Check priority.
	if e.Priority < c.MinPriority {
		contextutils.Logger(ctx).Infow(
			"priority not match",
			"event", e.UUID.String(),
			"alert_config", c.UUID.String(),
			"priority", e.Priority.String(),
			"need_priority", c.MinPriority.String(),
		)

		return false, nil
	}

	contextutils.Logger(ctx).Debugw(
		"priority matched",
		"event", e.UUID.String(),
		"alert_config", c.UUID.String(),
		"priority", e.Priority.String(),
		"need_priority", c.MinPriority.String(),
	)

	// Event matches alert config.
	if err := a.alertHitRepo.Create(ctx, c.UUID); err != nil {
		contextutils.Logger(ctx).Warnw(
			"cannot create alert hit",
			"error", err,
		)
		return false, fmt.Errorf("create alert hit: %w", err)
	}

	contextutils.Logger(ctx).Debugw(
		"created alert hit",
		"event", e.UUID.String(),
		"alert_config", c.UUID.String(),
	)

	rate, err := a.alertHitRepo.Get(ctx, c.UUID, c.Duration)
	if err != nil {
		contextutils.Logger(ctx).Warnw(
			"cannot get event rate",
			"error", err,
			"event", e.UUID,
			"event_message", e.Message,
			"duration", c.Duration,
			"alert_config", c.UUID,
		)

		return false, fmt.Errorf("get hits of alert: %w", err)
	}

	contextutils.Logger(ctx).Debugw(
		"got alert hit rate",
		"event", e.UUID.String(),
		"alert_config", c.UUID.String(),
		"rate", rate,
		"min_rate", c.MinRate,
	)

	return rate >= c.MinRate, nil
}

func (a *ProcessEventSyncManager) ProcessEvent(ctx context.Context, e *models.Event) error {
	contextutils.Logger(ctx).Infow(
		"processing event",
		"uuid", e.UUID,
		"message", e.Message,
		"priority", e.Priority.String(),
	)

	configs, err := a.alertConfigRepo.GetList(ctx)
	if err != nil {
		contextutils.Logger(ctx).Errorw("cannot get alert configs", "error", err)
		return fmt.Errorf("get alert configs: %w", err)
	}

	for _, c := range configs {
		needAlert, err := a.processAlert(ctx, e, c)
		if err != nil {
			contextutils.Logger(ctx).Errorw("cannot check event alert", "error", err)
			return fmt.Errorf("check event alert: %w", err)
		}

		if needAlert {
			contextutils.Logger(ctx).Infow(
				"found matching alert",
				"alert_uuid", c.UUID,
				"message_expression", c.MessageExpression,
				"comment", c.Comment,
			)

			err := a.alertCreationRepo.Create(ctx, &models.Alert{
				Message: c.Comment,
				Time:    time.Now().UTC(),
				Event:   e,
			})
			if err != nil {
				contextutils.Logger(ctx).Errorw("cannot create alert", "error", err)
				return fmt.Errorf("create alert: %w", err)
			}

			err = a.alertHitRepo.DeleteAll(ctx, c.UUID)
			if err != nil {
				contextutils.Logger(ctx).Errorw("cannot delete alert hits", "error", err)
				return fmt.Errorf("delete alert hits: %w", err)
			}
		}
	}

	return nil
}
