package managers

import (
	"context"
	"fmt"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/contextutils"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/ports"
)

type ProcessEventAsyncManager struct {
	alertProcessingQueue ports.AlertProcessingQueue
}

func NewProcessEventAsyncManager(apq ports.AlertProcessingQueue) *ProcessEventAsyncManager {
	return &ProcessEventAsyncManager{
		alertProcessingQueue: apq,
	}
}

func (a *ProcessEventAsyncManager) ProcessEvent(ctx context.Context, e *models.Event) error {
	contextutils.Logger(ctx).Infow(
		"processing event async",
		"uuid", e.UUID,
		"message", e.Message,
		"priority", e.Priority.String(),
	)

	if err := a.alertProcessingQueue.Push(ctx, e); err != nil {
		return fmt.Errorf("push event to process: %w", err)
	}

	return nil
}
