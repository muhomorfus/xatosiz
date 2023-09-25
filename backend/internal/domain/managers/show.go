package managers

import (
	"context"
	"fmt"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/contextutils"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/ports"
	"github.com/google/uuid"
)

type ShowManager struct {
	groupRepo ports.GroupRepository
	traceRepo ports.TraceRepository
	eventRepo ports.EventRepository
}

func NewShowManager(groupRepo ports.GroupRepository, traceRepo ports.TraceRepository, eventRepo ports.EventRepository) *ShowManager {
	return &ShowManager{
		groupRepo: groupRepo,
		traceRepo: traceRepo,
		eventRepo: eventRepo,
	}
}

func (m *ShowManager) GetGroupList(ctx context.Context, filters models.Filters) ([]*models.GroupPreview, []*models.GroupPreview, error) {
	contextutils.Logger(ctx).Infow("get group list")

	groups, err := m.groupRepo.GetList(ctx, filters)
	if err != nil {
		contextutils.Logger(ctx).Warnw("cant get group list", "error", err)

		return nil, nil, fmt.Errorf("get list of groups: %w", err)
	}

	activeGroups := make([]*models.GroupPreview, 0, len(groups))
	fixedGroups := make([]*models.GroupPreview, 0, len(groups))

	for _, g := range groups {
		if g.HasActiveEvents {
			activeGroups = append(activeGroups, g)
		} else {
			fixedGroups = append(fixedGroups, g)
		}
	}

	return activeGroups, fixedGroups, nil
}

func (m *ShowManager) GetGroup(ctx context.Context, id uuid.UUID) (*models.Group, error) {
	contextutils.Logger(ctx).Infow("get group", "uuid", id)

	group, err := m.groupRepo.Get(ctx, id)
	if err != nil {
		contextutils.Logger(ctx).Warnw("cant get group", "uuid", id, "error", err)

		return nil, fmt.Errorf("get group: %w", err)
	}

	return group, nil
}

func (m *ShowManager) GetEventList(ctx context.Context, fixed bool) ([]*models.Event, error) {
	contextutils.Logger(ctx).Infow("get event list")

	events, err := m.eventRepo.GetList(ctx, fixed)
	if err != nil {
		contextutils.Logger(ctx).Warnw("cant get event list", "error", err)

		return nil, fmt.Errorf("get event list: %w", err)
	}

	return events, nil
}

func (m *ShowManager) GetEvent(ctx context.Context, id uuid.UUID) (*models.Event, error) {
	contextutils.Logger(ctx).Infow("get event", "uuid", id)

	event, err := m.eventRepo.Get(ctx, id)
	if err != nil {
		contextutils.Logger(ctx).Warnw("cant get event", "uuid", id, "error", err)

		return nil, fmt.Errorf("get event: %w", err)
	}

	return event, nil
}

func (m *ShowManager) FixEvent(ctx context.Context, id uuid.UUID) error {
	contextutils.Logger(ctx).Infow("fix event", "uuid", id)

	event, err := m.eventRepo.Get(ctx, id)
	if err != nil {
		contextutils.Logger(ctx).Warnw("cant get event to fix", "uuid", id, "error", err)

		return fmt.Errorf("get event: %w", err)
	}

	event.Fixed = true

	if err := m.eventRepo.Update(ctx, event); err != nil {
		contextutils.Logger(ctx).Warnw("cant update event", "uuid", id, "error", err)

		return fmt.Errorf("update event: %w", err)
	}

	return nil
}
