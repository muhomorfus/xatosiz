package ports

import (
	"context"
	"time"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"github.com/google/uuid"
)

type GroupRepository interface {
	Create(ctx context.Context) (*models.Group, error)
	Get(ctx context.Context, id uuid.UUID) (*models.Group, error)
	GetList(ctx context.Context, filters models.Filters) ([]*models.GroupPreview, error)
}

type TraceRepository interface {
	Create(ctx context.Context, trace *models.Trace) (*models.Trace, error)
	Update(ctx context.Context, trace *models.Trace) error
	Get(ctx context.Context, id uuid.UUID) (*models.Trace, error)
}

type EventRepository interface {
	Create(ctx context.Context, event *models.Event) (*models.Event, error)
	Get(ctx context.Context, id uuid.UUID) (*models.Event, error)
	GetList(ctx context.Context, fixed bool) ([]*models.Event, error)
	Update(ctx context.Context, trace *models.Event) error
}

type AlertHitRepository interface {
	Get(ctx context.Context, id uuid.UUID, duration time.Duration) (int, error)
	Create(ctx context.Context, id uuid.UUID) error
	DeleteAll(ctx context.Context, id uuid.UUID) error
}

type AlertConfigRepository interface {
	Create(ctx context.Context, c *models.AlertConfig) (*models.AlertConfig, error)
	GetList(ctx context.Context) ([]*models.AlertConfig, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type AlertCreationRepository interface {
	Create(ctx context.Context, a *models.Alert) error
}

type AlertRepository interface {
	AlertCreationRepository
	GetList(ctx context.Context, solved bool) ([]models.Alert, error)
	MarkSolved(ctx context.Context, id uuid.UUID) error
}
