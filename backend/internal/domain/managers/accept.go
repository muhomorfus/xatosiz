package managers

import (
	"context"
	"fmt"
	"time"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/contextutils"
	domainErrors "git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/errors"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/ports"
	"github.com/google/uuid"
)

type AcceptManager struct {
	groupRepo    ports.GroupRepository
	traceRepo    ports.TraceRepository
	eventRepo    ports.EventRepository
	procEventMng ports.ProcessEventManager
}

func NewAcceptManager(groupRepo ports.GroupRepository, traceRepo ports.TraceRepository, eventRepo ports.EventRepository, procEventMng ports.ProcessEventManager) *AcceptManager {
	return &AcceptManager{
		groupRepo:    groupRepo,
		traceRepo:    traceRepo,
		eventRepo:    eventRepo,
		procEventMng: procEventMng,
	}
}

func (m *AcceptManager) StartTrace(ctx context.Context, trace models.TraceBaseInfo) (*models.Trace, error) {
	contextutils.Logger(ctx).Infow("start trace", "parent", trace.ParentUUID, "group", trace.GroupUUID, "title", trace.Title, "component", trace.Component)

	if trace.ParentUUID != nil {
		contextutils.Logger(ctx).Debugw("create trace with parent")

		parent, err := m.traceRepo.Get(ctx, *trace.ParentUUID)
		if err != nil {
			contextutils.Logger(ctx).Warnw("cant get parent trace", "parent", *trace.ParentUUID, "error", err)

			return nil, fmt.Errorf("get parent trace: %w", err)
		}

		if parent.Finished {
			contextutils.Logger(ctx).Errorw("trying to create child trace to finished")

			return nil, fmt.Errorf("check parent trace: %w", domainErrors.ErrTraceAlreadyFinished)
		}

		trace.GroupUUID = &parent.GroupUUID
	}

	if trace.GroupUUID == nil {
		contextutils.Logger(ctx).Debugw("create trace without group and without parent")

		group, err := m.groupRepo.Create(ctx)
		if err != nil {
			contextutils.Logger(ctx).Errorw("cant create group", "error", err)

			return nil, fmt.Errorf("create group: %w", err)
		}

		trace.GroupUUID = &group.UUID
	} else {
		contextutils.Logger(ctx).Debugw("create trace with group")

		if _, err := m.groupRepo.Get(ctx, *trace.GroupUUID); err != nil {
			contextutils.Logger(ctx).Warnw("cant get group", "uuid", *trace.GroupUUID, "error", err)

			return nil, fmt.Errorf("get group: %w", err)
		}
	}

	newTrace := &models.Trace{
		GroupUUID:  *trace.GroupUUID,
		ParentUUID: trace.ParentUUID,
		Start:      time.Now().UTC(),
		Title:      trace.Title,
		Component:  trace.Component,
	}

	newTrace, err := m.traceRepo.Create(ctx, newTrace)
	if err != nil {
		contextutils.Logger(ctx).Errorw("cant create trace", "error", err)

		return nil, fmt.Errorf("create trace: %w", err)
	}

	contextutils.Logger(ctx).Infow("started trace", "uuid", newTrace.UUID)

	return newTrace, nil
}

func (m *AcceptManager) EndTrace(ctx context.Context, id uuid.UUID) error {
	contextutils.Logger(ctx).Infow("end trace", "uuid", id)

	trace, err := m.traceRepo.Get(ctx, id)
	if err != nil {
		contextutils.Logger(ctx).Warnw("cant get trace", "uuid", id, "error", err)

		return fmt.Errorf("get trace: %w", err)
	}

	if trace.Finished {
		contextutils.Logger(ctx).Errorw("trying to finish trace already finished")

		return fmt.Errorf("check trace status: %w", domainErrors.ErrTraceAlreadyFinished)
	}

	trace.Finished = true
	trace.End = time.Now().UTC()

	if err := m.traceRepo.Update(ctx, trace); err != nil {
		contextutils.Logger(ctx).Warnw("cant update trace", "uuid", trace.UUID, "error", err)

		return fmt.Errorf("update trace: %w", err)
	}

	contextutils.Logger(ctx).Infow("ended trace", "uuid", id)

	return nil
}

func (m *AcceptManager) CreateEvent(ctx context.Context, event models.EventBaseInfo) (*models.Event, error) {
	contextutils.Logger(ctx).Infow("create event", "trace", event.TraceUUID, "group", event.GroupUUID, "message", event.Message)

	var trace *models.Trace
	var err error

	if event.TraceUUID == nil {
		contextutils.Logger(ctx).Debugw("create event without trace")

		trace, err = m.StartTrace(ctx, models.TraceBaseInfo{
			GroupUUID:  event.GroupUUID,
			ParentUUID: nil,
			Title:      event.Message,
			Component:  event.Component,
		})

		if err != nil {
			contextutils.Logger(ctx).Warnw("cant start trace", "group", *event.GroupUUID, "error", err)

			return nil, fmt.Errorf("start trace: %w", err)
		}

		event.TraceUUID = &trace.UUID
	} else {
		contextutils.Logger(ctx).Debugw("create event with trace")

		trace, err = m.traceRepo.Get(ctx, *event.TraceUUID)
		if err != nil {
			contextutils.Logger(ctx).Warnw("cant get trace", "trace", *event.TraceUUID, "error", err)

			return nil, fmt.Errorf("get trace: %w", err)
		}
	}

	if trace.Finished {
		contextutils.Logger(ctx).Errorw("trying to create event to finished trace")

		return nil, fmt.Errorf("check trace: %w", domainErrors.ErrTraceAlreadyFinished)
	}

	newEvent := &models.Event{
		TraceUUID: *event.TraceUUID,
		GroupUUID: trace.GroupUUID,
		Message:   event.Message,
		Priority:  event.Priority,
		Payload:   event.Payload,
		Time:      time.Now().UTC(),
	}
	newEvent, err = m.eventRepo.Create(ctx, newEvent)
	if err != nil {
		contextutils.Logger(ctx).Errorw("cant create event", "error", err)

		return nil, fmt.Errorf("create event: %w", err)
	}

	contextutils.Logger(ctx).Infow("created event", "uuid", newEvent.UUID)

	if err := m.procEventMng.ProcessEvent(ctx, newEvent); err != nil {
		contextutils.Logger(ctx).Errorw("cant process event (check alerting)", "error", err)

		return nil, fmt.Errorf("check alert: %w", err)
	}

	return newEvent, nil
}

func (m *AcceptManager) CreateGroup(ctx context.Context) (*models.Group, error) {
	contextutils.Logger(ctx).Infow("create group")

	group, err := m.groupRepo.Create(ctx)
	if err != nil {
		contextutils.Logger(ctx).Errorw("cant create group", "error", err)

		return nil, fmt.Errorf("create group: %w", err)
	}

	contextutils.Logger(ctx).Infow("created group", "uuid", group.UUID)

	return group, nil
}

func toPtr[T any](v T) *T {
	return &v
}
