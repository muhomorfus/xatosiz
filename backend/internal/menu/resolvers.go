package menu

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"github.com/dixonwille/wmenu/v5"
	"github.com/google/uuid"
)

func output(v interface{}) {
	bytes, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(bytes))
}

func (m *Menu) createGroup(opt wmenu.Opt) error {
	ctx := opt.Value.(context.Context)

	g, err := m.accept.CreateGroup(ctx)
	if err != nil {
		return fmt.Errorf("create group: %w", err)
	}

	fmt.Println("Group:")
	output(g)

	return nil
}

func (m *Menu) startTrace(opt wmenu.Opt) error {
	ctx := opt.Value.(context.Context)

	var groupUUID, parentUUID *uuid.UUID
	var groupUUIDString, parentUUIDString string
	var title, component string

	fmt.Print("Group UUID: ")
	if _, err := fmt.Scanln(&groupUUIDString); err == nil {
		id, err := uuid.Parse(groupUUIDString)
		if err != nil {
			return fmt.Errorf("parse group uuid: %w", err)
		}

		groupUUID = &id
	}

	fmt.Print("Parent UUID: ")
	if _, err := fmt.Scanln(&parentUUIDString); err == nil {
		id, err := uuid.Parse(parentUUIDString)
		if err != nil {
			return fmt.Errorf("parse parent uuid: %w", err)
		}

		parentUUID = &id
	}

	fmt.Print("Title: ")
	if _, err := fmt.Scanln(&title); err != nil {
		return fmt.Errorf("scan title: %w", err)
	}

	fmt.Print("Component: ")
	if _, err := fmt.Scanln(&component); err != nil {
		return fmt.Errorf("scan component: %w", err)
	}

	trace, err := m.accept.StartTrace(ctx, models.TraceBaseInfo{
		GroupUUID:  groupUUID,
		ParentUUID: parentUUID,
		Title:      title,
		Component:  component,
	})
	if err != nil {
		return fmt.Errorf("start trace: %w", err)
	}

	fmt.Println("Trace:")
	output(trace)

	return nil
}

func (m *Menu) endTrace(opt wmenu.Opt) error {
	ctx := opt.Value.(context.Context)

	fmt.Print("UUID: ")
	var idString string
	if _, err := fmt.Scanln(&idString); err != nil {
		return fmt.Errorf("scan uuid: %w", err)
	}

	id, err := uuid.Parse(idString)
	if err != nil {
		return fmt.Errorf("parse trace uuid: %w", err)
	}

	if err := m.accept.EndTrace(ctx, id); err != nil {
		return fmt.Errorf("end trace: %w", err)
	}

	return nil
}

func (m *Menu) sendEvent(opt wmenu.Opt) error {
	ctx := opt.Value.(context.Context)

	var groupUUID, traceUUID *uuid.UUID
	var groupUUIDString, traceUUIDString string
	var message, component string
	var priority int // 0-3
	var payloadString string
	var payload map[string]string

	fmt.Print("Group UUID: ")
	if _, err := fmt.Scanln(&groupUUIDString); err == nil {
		id, err := uuid.Parse(groupUUIDString)
		if err != nil {
			return fmt.Errorf("parse group uuid: %w", err)
		}

		groupUUID = &id
	}

	fmt.Print("Trace UUID: ")
	if _, err := fmt.Scanln(&traceUUIDString); err == nil {
		id, err := uuid.Parse(traceUUIDString)
		if err != nil {
			return fmt.Errorf("parse trace uuid: %w", err)
		}

		traceUUID = &id
	}

	fmt.Print("Message: ")
	if _, err := fmt.Scanln(&message); err != nil {
		return fmt.Errorf("scan title: %w", err)
	}

	fmt.Print("Component: ")
	if _, err := fmt.Scanln(&component); err != nil {
		return fmt.Errorf("scan component: %w", err)
	}

	fmt.Print("Priority (0-3): ")
	if _, err := fmt.Scanln(&priority); err != nil {
		return fmt.Errorf("scan priority: %w", err)
	}

	fmt.Print("Payload (JSON without spaces): ")
	if _, err := fmt.Scanln(&payloadString); err != nil {
		return fmt.Errorf("scan priority: %w", err)
	}

	if err := json.Unmarshal([]byte(payloadString), &payload); err != nil {
		return fmt.Errorf("parse json: %w", err)
	}

	event, err := m.accept.CreateEvent(ctx, models.EventBaseInfo{
		GroupUUID: groupUUID,
		TraceUUID: traceUUID,
		Message:   message,
		Priority:  models.Priority(priority),
		Payload:   payload,
		Component: component,
	})
	if err != nil {
		return fmt.Errorf("create event: %w", err)
	}

	fmt.Println("Event:")
	output(event)

	return nil
}

func (m *Menu) groupList(opt wmenu.Opt) error {
	ctx := opt.Value.(context.Context)

	active, fixed, err := m.show.GetGroupList(ctx, models.Filters{})
	if err != nil {
		return fmt.Errorf("get groups: %w", err)
	}

	fmt.Println("Active groups:")
	output(active)

	fmt.Println("Fixed groups:")
	output(fixed)

	return nil
}

func (m *Menu) eventList(opt wmenu.Opt) error {
	ctx := opt.Value.(context.Context)

	events, err := m.show.GetEventList(ctx, false)
	if err != nil {
		return fmt.Errorf("get event list: %w", err)
	}

	fmt.Println("Events:")
	output(events)

	return nil
}

func (m *Menu) showEvent(opt wmenu.Opt) error {
	ctx := opt.Value.(context.Context)

	fmt.Print("UUID: ")
	var idString string
	if _, err := fmt.Scanln(&idString); err != nil {
		return fmt.Errorf("scan uuid: %w", err)
	}

	id, err := uuid.Parse(idString)
	if err != nil {
		return fmt.Errorf("parse event uuid: %w", err)
	}

	event, err := m.show.GetEvent(ctx, id)
	if err != nil {
		return fmt.Errorf("get event: %w", err)
	}

	fmt.Println("Event:")
	output(event)

	return nil
}

func (m *Menu) fixEvent(opt wmenu.Opt) error {
	ctx := opt.Value.(context.Context)

	fmt.Print("UUID: ")
	var idString string
	if _, err := fmt.Scanln(&idString); err != nil {
		return fmt.Errorf("scan uuid: %w", err)
	}

	id, err := uuid.Parse(idString)
	if err != nil {
		return fmt.Errorf("parse event uuid: %w", err)
	}

	if err := m.show.FixEvent(ctx, id); err != nil {
		return fmt.Errorf("fix event: %w", err)
	}

	return nil
}

func (m *Menu) addAlertConfig(opt wmenu.Opt) error {
	ctx := opt.Value.(context.Context)

	var messageExpression, minPriority string
	var duration string
	var minRate int
	var comment string

	fmt.Print("Message regular expression: ")
	if _, err := fmt.Scanln(&messageExpression); err != nil {
		return fmt.Errorf("scan messageExpression: %w", err)
	}

	fmt.Print("Minimum priority: ")
	if _, err := fmt.Scanln(&minPriority); err != nil {
		return fmt.Errorf("scan minPriority: %w", err)
	}

	fmt.Print("Duration: ")
	if _, err := fmt.Scanln(&duration); err != nil {
		return fmt.Errorf("scan duration: %w", err)
	}
	parsedDuration, err := time.ParseDuration(duration)
	if err != nil {
		return fmt.Errorf("parse daration: %w", err)
	}

	fmt.Print("Min rate: ")
	if _, err := fmt.Scanln(&minRate); err != nil {
		return fmt.Errorf("scan minRate: %w", err)
	}

	fmt.Print("Comment: ")
	if _, err := fmt.Scanln(&comment); err != nil {
		return fmt.Errorf("scan comment: %w", err)
	}

	cfg, err := m.alertCfg.CreateAlertConfig(ctx, &models.AlertConfig{
		MessageExpression: messageExpression,
		MinPriority:       models.PriorityFromString(minPriority),
		Duration:          parsedDuration,
		MinRate:           minRate,
		Comment:           comment,
	})
	if err != nil {
		return fmt.Errorf("create alert config: %w", err)
	}

	output(cfg)

	return nil
}

func (m *Menu) alertConfigList(opt wmenu.Opt) error {
	ctx := opt.Value.(context.Context)

	list, err := m.alertCfg.GetAlertConfigList(ctx)
	if err != nil {
		return fmt.Errorf("get alert config list: %w", err)
	}

	output(list)

	return nil
}

func (m *Menu) deleteAlertConfig(opt wmenu.Opt) error {
	ctx := opt.Value.(context.Context)

	fmt.Print("UUID: ")
	var idString string
	if _, err := fmt.Scanln(&idString); err != nil {
		return fmt.Errorf("scan uuid: %w", err)
	}

	id, err := uuid.Parse(idString)
	if err != nil {
		return fmt.Errorf("parse cfg uuid: %w", err)
	}

	if err := m.alertCfg.DeleteAlertConfig(ctx, id); err != nil {
		return fmt.Errorf("delete alert config: %w", err)
	}

	return nil
}

func (m *Menu) alertList(opt wmenu.Opt) error {
	ctx := opt.Value.(context.Context)

	list, err := m.alert.GetAlertList(ctx, false)
	if err != nil {
		return fmt.Errorf("get alerts: %w", err)
	}

	output(list)

	return nil
}
