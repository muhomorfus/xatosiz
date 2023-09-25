package ports

import (
	"context"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
)

type AlertProcessingQueue interface {
	Push(ctx context.Context, e *models.Event) error
}
