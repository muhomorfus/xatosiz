package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"github.com/google/uuid"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/openapi"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/ports"
)

type Server struct {
	accept   ports.AcceptManager
	show     ports.ShowManager
	alert    ports.AlertManager
	alertCfg ports.AlertConfigManager
}

func New(accept ports.AcceptManager, show ports.ShowManager, alert ports.AlertManager, ac ports.AlertConfigManager) *Server {
	return &Server{
		accept:   accept,
		show:     show,
		alert:    alert,
		alertCfg: ac,
	}
}

func (s *Server) CreateGroup(ctx context.Context) (openapi.ImplResponse, error) {
	g, err := s.accept.CreateGroup(ctx)
	if err != nil {
		return openapi.ImplResponse{
			Code: http.StatusInternalServerError,
			Body: openapi.CreateGroupResponse{
				Uuid: g.UUID.String(),
			},
		}, fmt.Errorf("create group: %w", err)
	}

	return openapi.ImplResponse{
		Code: http.StatusCreated,
		Body: openapi.CreateGroupResponse{
			Uuid: g.UUID.String(),
		},
	}, nil
}

func (s *Server) EndTrace(ctx context.Context, uuidStr string) (openapi.ImplResponse, error) {
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		return toError(err, "Cannot parse UUID.")
	}

	if err := s.accept.EndTrace(ctx, id); err != nil {
		return toError(err, "Cannot end trace.")
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
	}, nil
}

func (s *Server) FixEvent(ctx context.Context, uuidStr string) (openapi.ImplResponse, error) {
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		return toError(err, "Cannot parse UUID.")
	}

	if err := s.show.FixEvent(ctx, id); err != nil {
		return toError(err, "Cannot fix event.")
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
	}, nil
}

func (s *Server) GetEvent(ctx context.Context, uuidStr string) (openapi.ImplResponse, error) {
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		return toError(err, "Cannot parse UUID.")
	}

	event, err := s.show.GetEvent(ctx, id)
	if err != nil {
		return toError(err, "Cannot get event.")
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: toEvent(event),
	}, nil
}

func (s *Server) GetEventList(ctx context.Context) (openapi.ImplResponse, error) {
	events, err := s.show.GetEventList(ctx, false)
	if err != nil {
		return toError(err, "Cannot get events.")
	}

	result := make([]openapi.Event, len(events))
	for i := range result {
		result[i] = toEvent(events[i])
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.GetEventListResponse{Items: result},
	}, nil
}

func (s *Server) GetGroupList(ctx context.Context, filters openapi.Filters) (openapi.ImplResponse, error) {
	active, fixed, err := s.show.GetGroupList(ctx, models.Filters{
		Limit:     filters.Limit,
		Component: filters.Component,
	})
	if err != nil {
		return toError(err, "Cannot get groups.")
	}

	resultActive := make([]openapi.GroupPreview, len(active))
	for i := range resultActive {
		resultActive[i] = toGroupPreview(active[i])
	}

	resultFixed := make([]openapi.GroupPreview, len(fixed))
	for i := range resultFixed {
		resultFixed[i] = toGroupPreview(fixed[i])
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.GetGroupListResponse{
			Active: resultActive,
			Fixed:  resultFixed,
		},
	}, nil
}

func (s *Server) SendEvent(ctx context.Context, req openapi.SendEventRequest) (openapi.ImplResponse, error) {
	var groupID, traceID *uuid.UUID

	if req.GroupUuid != "" {
		id, err := uuid.Parse(req.GroupUuid)
		if err != nil {
			return toError(err, "Cannot parse group UUID.")
		}

		groupID = new(uuid.UUID)
		*groupID = id
	}

	if req.TraceUuid != "" {
		id, err := uuid.Parse(req.TraceUuid)
		if err != nil {
			return toError(err, "Cannot parse group UUID.")
		}

		traceID = new(uuid.UUID)
		*traceID = id
	}

	event, err := s.accept.CreateEvent(ctx, models.EventBaseInfo{
		GroupUUID: groupID,
		TraceUUID: traceID,
		Message:   req.Message,
		Priority:  fromPriority(req.Priority),
		Payload:   req.Payload,
		Component: req.Component,
	})
	if err != nil {
		return toError(err, "Cannot create event.")
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: toEvent(event),
	}, nil
}

func (s *Server) StartTrace(ctx context.Context, req openapi.StartTraceRequest) (openapi.ImplResponse, error) {
	var groupID, parentID *uuid.UUID

	if req.GroupUuid != "" {
		id, err := uuid.Parse(req.GroupUuid)
		if err != nil {
			return toError(err, "Cannot parse group UUID.")
		}

		groupID = new(uuid.UUID)
		*groupID = id
	}

	if req.ParentUuid != "" {
		id, err := uuid.Parse(req.ParentUuid)
		if err != nil {
			return toError(err, "Cannot parse group UUID.")
		}

		parentID = new(uuid.UUID)
		*parentID = id
	}

	trace, err := s.accept.StartTrace(ctx, models.TraceBaseInfo{
		GroupUUID:  groupID,
		ParentUUID: parentID,
		Title:      req.Title,
		Component:  req.Component,
	})
	if err != nil {
		return toError(err, "Cannot start trace.")
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: toStartTraceResponse(trace),
	}, nil
}

func (s *Server) GetGroup(ctx context.Context, idString string) (openapi.ImplResponse, error) {
	id, err := uuid.Parse(idString)
	if err != nil {
		return toError(err, "Cannot parse group UUID.")
	}

	group, err := s.show.GetGroup(ctx, id)
	if err != nil {
		return toError(err, "Cannot get group info.")
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: toGroup(group),
	}, nil
}

func (s *Server) CreateAlertConfig(ctx context.Context, cfg openapi.AlertConfig) (openapi.ImplResponse, error) {
	duration, err := time.ParseDuration(cfg.Duration)
	if err != nil {
		return toError(err, "Cannot parse duration.")
	}

	alert, err := s.alertCfg.CreateAlertConfig(ctx, &models.AlertConfig{
		MessageExpression: cfg.MessageExpression,
		MinPriority:       models.PriorityFromString(cfg.MinPriority),
		Duration:          duration,
		MinRate:           int(cfg.MinRate),
		Comment:           cfg.Comment,
	})
	if err != nil {
		return toError(err, "Cannot create alert config.")
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.AlertConfig{
			Uuid:              alert.UUID.String(),
			MessageExpression: alert.MessageExpression,
			MinPriority:       alert.MinPriority.String(),
			Duration:          alert.Duration.String(),
			MinRate:           int32(alert.MinRate),
			Comment:           alert.Comment,
		},
	}, nil
}

func (s *Server) DeleteAlertConfig(ctx context.Context, idStr string) (openapi.ImplResponse, error) {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return toError(err, "Cannot parse group UUID.")
	}

	if err := s.alertCfg.DeleteAlertConfig(ctx, id); err != nil {
		return toError(err, "Cannot delete alert config.")
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
	}, nil
}

func (s *Server) FixAlert(ctx context.Context, idStr string) (openapi.ImplResponse, error) {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return toError(err, "Cannot parse group UUID.")
	}

	if err := s.alert.SolveAlert(ctx, id); err != nil {
		return toError(err, "Cannot fix alert.")
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
	}, nil
}

func (s *Server) GetAlertConfigList(ctx context.Context) (openapi.ImplResponse, error) {
	list, err := s.alertCfg.GetAlertConfigList(ctx)
	if err != nil {
		return toError(err, "Cannot get alert list.")
	}

	result := make([]openapi.AlertConfig, len(list))
	for i := range result {
		result[i] = openapi.AlertConfig{
			Uuid:              list[i].UUID.String(),
			MessageExpression: list[i].MessageExpression,
			MinPriority:       list[i].MinPriority.String(),
			Duration:          list[i].Duration.String(),
			MinRate:           int32(list[i].MinRate),
			Comment:           list[i].Comment,
		}
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.GetAlertConfigListResponse{Items: result},
	}, nil
}

func (s *Server) GetAlertList(ctx context.Context) (openapi.ImplResponse, error) {
	alerts, err := s.alert.GetAlertList(ctx, false)
	if err != nil {
		return toError(err, "Cannot get alerts.")
	}

	result := make([]openapi.GetAlertListResponseItemsInner, len(alerts))

	for i := range result {
		result[i] = openapi.GetAlertListResponseItemsInner{
			Uuid:    alerts[i].UUID.String(),
			Message: alerts[i].Message,
			Time:    alerts[i].Time.Format(time.RFC3339Nano),
			Event:   toEvent(alerts[i].Event),
		}
	}

	return openapi.ImplResponse{
		Code: http.StatusOK,
		Body: openapi.GetAlertListResponse{Items: result},
	}, nil
}
