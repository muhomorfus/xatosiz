package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func toGroup(ctx context.Context, db *gorm.DB, g Group) (*models.Group, error) {
	var traces []Trace
	if err := db.WithContext(ctx).Table("trace").Where("group_uuid = ? and parent_uuid is null", g.UUID).Find(&traces).Error; err != nil {
		return nil, fmt.Errorf("get group traces: %w", err)
	}

	result := &models.Group{
		UUID: g.UUID,
	}
	result.Traces = make([]models.Trace, len(traces))
	for i, c := range traces {
		var err error

		if result.Traces[i], err = toTrace(ctx, db, c); err != nil {
			return nil, fmt.Errorf("convert group trace: %w", err)
		}
	}

	return result, nil
}

func toGroupPreview(ctx context.Context, gp GroupPreview, db *gorm.DB) (*models.GroupPreview, error) {
	var components []Component

	query := `
		select
			component as name,
			count(*) as quantity
		from trace
		where group_uuid = ?
		group by component;
	`

	if err := db.WithContext(ctx).Raw(query, gp.UUID).Find(&components).Error; err != nil {
		return nil, fmt.Errorf("get group components: %w", err)
	}

	var events []Event

	if err := db.WithContext(ctx).Table("event").Where("group_uuid = ?", gp.UUID).Find(&events).Error; err != nil {
		return nil, fmt.Errorf("get group events: %w", err)
	}

	resultComponents := make([]models.Component, len(components))
	for i := range components {
		resultComponents[i] = models.Component{
			Name:     components[i].Name,
			Quantity: components[i].Quantity,
		}
	}

	resultEvents := make([]models.Event, len(events))
	for i := range events {
		var err error
		if resultEvents[i], err = toEvent(events[i]); err != nil {
			return nil, fmt.Errorf("convert event: %w", err)
		}
	}

	result := &models.GroupPreview{
		UUID:            gp.UUID,
		Title:           gp.Title,
		HasActiveEvents: !gp.Fixed,
		Events:          resultEvents,
		Components:      resultComponents,
		Start:           gp.Start,
		End:             gp.End,
		Duration:        gp.End.Sub(gp.Start),
	}

	return result, nil
}

func toTrace(ctx context.Context, db *gorm.DB, t Trace) (models.Trace, error) {
	end := time.Now().UTC()
	if t.TimeEnd != nil {
		end = *t.TimeEnd
	}

	result := models.Trace{
		UUID:       t.UUID,
		GroupUUID:  t.GroupUUID,
		ParentUUID: t.ParentUUID,
		Start:      t.TimeStart,
		End:        end,
		Title:      t.Title,
		Component:  t.Component,
		Finished:   t.TimeEnd != nil,
	}

	var err error
	if result.Events, err = getEvents(ctx, db, result.UUID); err != nil {
		return models.Trace{}, fmt.Errorf("get events: %w", err)
	}

	var children []Trace
	if err := db.WithContext(ctx).Table("trace").Where("parent_uuid = ?", t.UUID).Find(&children).Error; err != nil {
		return models.Trace{}, fmt.Errorf("get child traces: %w", err)
	}

	result.Children = make([]models.Trace, len(children))
	for i, c := range children {
		if result.Children[i], err = toTrace(ctx, db, c); err != nil {
			return models.Trace{}, fmt.Errorf("convert child trace: %w", err)
		}
	}

	return result, nil
}

func getEvents(ctx context.Context, db *gorm.DB, traceID uuid.UUID) ([]models.Event, error) {
	var events []Event

	if err := db.WithContext(ctx).Table("event").Where("trace_uuid = ?", traceID).Find(&events).Error; err != nil {
		return nil, fmt.Errorf("get events from postgres: %w", err)
	}

	result := make([]models.Event, len(events))
	for i, e := range events {
		var err error
		if result[i], err = toEvent(e); err != nil {
			return nil, fmt.Errorf("eonvert event: %w", err)
		}
	}

	return result, nil
}

func toEvent(e Event) (models.Event, error) {
	res := models.Event{
		UUID:      e.UUID,
		TraceUUID: e.TraceUUID,
		Message:   e.Message,
		Priority:  priorityToDomain(e.Priority),
		Fixed:     e.Fixed,
		Time:      e.Time,
	}

	if e.Payload != nil {
		if err := json.Unmarshal([]byte(*(e.Payload)), &res.Payload); err != nil {
			return models.Event{}, fmt.Errorf("unmarshal payload: %w", err)
		}
	}

	return res, nil
}
