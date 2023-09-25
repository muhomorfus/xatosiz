package ports

import (
	"context"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"github.com/google/uuid"
)

type AcceptManager interface {
	StartTrace(ctx context.Context, trace models.TraceBaseInfo) (*models.Trace, error)
	EndTrace(ctx context.Context, id uuid.UUID) error
	CreateEvent(ctx context.Context, event models.EventBaseInfo) (*models.Event, error)
	CreateGroup(ctx context.Context) (*models.Group, error)
}

type ShowManager interface {
	GetGroupList(ctx context.Context, filters models.Filters) ([]*models.GroupPreview, []*models.GroupPreview, error)
	GetGroup(ctx context.Context, id uuid.UUID) (*models.Group, error)
	GetEventList(ctx context.Context, fixed bool) ([]*models.Event, error)
	GetEvent(ctx context.Context, id uuid.UUID) (*models.Event, error)
	FixEvent(ctx context.Context, id uuid.UUID) error
}

type AlertSaveManager interface {
	Save(ctx context.Context, alert *models.Alert) error
}

type AlertManager interface {
	GetAlertList(ctx context.Context, solved bool) ([]models.Alert, error)
	SolveAlert(ctx context.Context, id uuid.UUID) error
}

type AlertConfigManager interface {
	CreateAlertConfig(ctx context.Context, c *models.AlertConfig) (*models.AlertConfig, error)
	GetAlertConfigList(ctx context.Context) ([]*models.AlertConfig, error)
	DeleteAlertConfig(ctx context.Context, id uuid.UUID) error
}

type ProcessEventManager interface {
	ProcessEvent(ctx context.Context, e *models.Event) error
}
