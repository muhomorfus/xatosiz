package tracing

import (
	"context"
	"fmt"
	"git.iu7.bmstu.ru/keo20u511/ppo/backend/internal/tracing/openapi"
	"github.com/AlekSi/pointer"
)

type Trace struct {
	id      string
	groupID string
}

func StartTrace(ctx context.Context, title string) (*Trace, context.Context) {
	var groupUUID, parentUUID *string

	groupID, ok := groupFromContext(ctx)
	if ok {
		groupUUID = &groupID
	}

	parentID, ok := traceFromContext(ctx)
	if ok {
		parentUUID = &parentID
	}

	logger.Debugw(
		"starting trace",
		"group_uuid", pointer.Get(groupUUID),
		"parent_uuid", pointer.Get(parentUUID),
		"title", title,
		"component", serviceName,
	)

	trace, _, err := api.DefaultApi.StartTrace(ctx).StartTraceRequest(openapi.StartTraceRequest{
		GroupUuid:  groupUUID,
		ParentUuid: parentUUID,
		Title:      title,
		Component:  serviceName,
	}).Execute()
	if err != nil {
		logger.Errorw(
			"cant start trace",
			"error", err,
			"group_uuid", pointer.Get(groupUUID),
			"parent_uuid", pointer.Get(parentUUID),
			"title", title,
			"component", serviceName,
		)
		return nil, ctx
	}

	logger.Debugw(
		"started",
		"uuid", trace.Uuid,
		"group_uuid", pointer.Get(trace.GroupUuid),
		"parent_uuid", pointer.Get(trace.ParentUuid),
		"title", trace.Title,
		"component", trace.Component,
	)

	return &Trace{
		id:      trace.Uuid,
		groupID: pointer.GetString(trace.GroupUuid),
	}, GroupToContext(TraceToContext(ctx, trace.Uuid), pointer.Get(trace.GroupUuid))
}

func (t *Trace) End() {
	if t == nil {
		logger.Warnw("try to end empty trace")
		return
	}

	logger.Debugw("ending trace", "trace_uuid", t.id)

	_, err := api.DefaultApi.EndTrace(context.Background(), t.id).Execute()
	if err != nil {
		logger.Warnw("cant end trace", "error", err, "trace_uuid", t.id)
	}
}

type Priority string

var (
	Info    Priority = "info"
	Warning Priority = "warning"
	Error   Priority = "error"
	Fatal   Priority = "fatal"
)

type Event struct {
	Message  string
	Priority Priority
	Payload  map[string]any
}

func (t *Trace) SendEvent(e *Event) {
	if t == nil {
		logger.Errorw("try send event to empty trace")
		return
	}

	logger.Debugw(
		"ending event",
		"group_uuid", t.groupID,
		"parent_uuid", t.id,
		"message", e.Message,
		"component", serviceName,
		"priority", e.Priority,
		"payload", e.Payload,
	)

	var payload map[string]string
	if e.Payload != nil {
		payload = make(map[string]string, len(e.Payload))

		for k, v := range e.Payload {
			payload[k] = fmt.Sprint(v)
		}
	}

	_, _, err := api.DefaultApi.SendEvent(context.Background()).SendEventRequest(openapi.SendEventRequest{
		GroupUuid: &t.groupID,
		TraceUuid: &t.id,
		Message:   e.Message,
		Component: serviceName,
		Priority:  string(e.Priority),
		Payload:   &payload,
	}).Execute()
	if err != nil {
		logger.Errorw(
			"cant send event",
			"error", err,
			"group_uuid", t.groupID,
			"parent_uuid", t.id,
			"message", e.Message,
			"component", serviceName,
			"priority", e.Priority,
			"payload", e.Payload,
		)
	}
}
