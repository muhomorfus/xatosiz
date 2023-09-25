package server

import (
	"errors"
	"net/http"
	"time"

	domainErrors "git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/errors"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/openapi"
)

func toError(err error, defaultComment string) (openapi.ImplResponse, error) {
	switch {
	case errors.Is(err, domainErrors.ErrTraceAlreadyFinished):
		return openapi.ImplResponse{
			Code: http.StatusBadRequest,
			Body: openapi.Error{
				Error:   err.Error(),
				Comment: "Trace already finished.",
			},
		}, nil
	default:
		return openapi.ImplResponse{
			Code: http.StatusInternalServerError,
			Body: openapi.Error{
				Error:   err.Error(),
				Comment: defaultComment,
			},
		}, nil
	}
}

func toEvent(event *models.Event) openapi.Event {
	return openapi.Event{
		Uuid:      event.UUID.String(),
		TraceUuid: event.TraceUUID.String(),
		Message:   event.Message,
		Priority:  event.Priority.String(),
		Fixed:     event.Fixed,
		Time:      event.Time.Format(time.RFC3339Nano),
		Payload:   event.Payload,
	}
}

func fromPriority(p string) models.Priority {
	switch p {
	case "info":
		return models.PriorityInfo
	case "warning":
		return models.PriorityWarning
	case "error":
		return models.PriorityError
	case "fatal":
		return models.PriorityFatal
	default:
		return models.PriorityInfo
	}
}

func toTrace(t *models.Trace) openapi.Trace {
	children := make([]openapi.Trace, len(t.Children))
	for i := range children {
		children[i] = toTrace(&t.Children[i])
	}

	events := make([]openapi.Event, len(t.Events))
	for i := range events {
		events[i] = toEvent(&t.Events[i])
	}

	var parentID string
	if t.ParentUUID != nil {
		parentID = t.ParentUUID.String()
	}

	return openapi.Trace{
		Uuid:       t.UUID.String(),
		GroupUuid:  t.GroupUUID.String(),
		ParentUuid: parentID,
		Start:      t.Start.Format(time.RFC3339Nano),
		End:        t.End.Format(time.RFC3339Nano),
		Title:      t.Title,
		Component:  t.Component,
		Finished:   t.Finished,
		Children:   children,
		Events:     events,
	}
}

func toStartTraceResponse(t *models.Trace) openapi.StartTraceResponse {
	parentUUID := ""
	if t.ParentUUID != nil {
		parentUUID = t.ParentUUID.String()
	}

	return openapi.StartTraceResponse{
		Uuid:       t.UUID.String(),
		GroupUuid:  t.GroupUUID.String(),
		ParentUuid: parentUUID,
		Start:      t.Start.Format(time.RFC3339Nano),
		End:        t.End.Format(time.RFC3339Nano),
		Title:      t.Title,
		Component:  t.Component,
		Finished:   t.Finished,
	}
}

func toGroup(g *models.Group) openapi.Group {
	traces := make([]openapi.Trace, len(g.Traces))
	for i := range traces {
		traces[i] = toTrace(&g.Traces[i])
	}

	return openapi.Group{
		Uuid:   g.UUID.String(),
		Traces: traces,
	}
}

func toGroupPreview(g *models.GroupPreview) openapi.GroupPreview {
	components := make([]openapi.Component, len(g.Components))
	for i := range components {
		components[i] = openapi.Component{
			Name:     g.Components[i].Name,
			Quantity: int32(g.Components[i].Quantity),
		}
	}

	events := make([]openapi.Event, len(g.Events))
	for i := range events {
		events[i] = toEvent(&g.Events[i])
	}

	return openapi.GroupPreview{
		Uuid:       g.UUID.String(),
		Title:      g.Title,
		Duration:   g.Duration.String(),
		Start:      g.Start.Format(time.RFC3339Nano),
		End:        g.End.Format(time.RFC3339Nano),
		Events:     events,
		Components: components,
	}
}
