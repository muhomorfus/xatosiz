package managers

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/mocks"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/ports"
	"github.com/gojuno/minimock/v3"
	"github.com/google/uuid"
)

func TestShowManager_GetGroupList(t *testing.T) {
	mc := minimock.NewController(t)

	groupUUIDs := []uuid.UUID{
		uuid.New(),
		uuid.New(),
		uuid.New(),
	}

	type fields struct {
		groupRepo ports.GroupRepository
		traceRepo ports.TraceRepository
		eventRepo ports.EventRepository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.GroupPreview
		want1   []*models.GroupPreview
		wantErr bool
	}{
		{
			name: "success get group list",
			fields: fields{
				groupRepo: mocks.NewGroupRepositoryMock(mc).GetListMock.Return([]*models.GroupPreview{
					{
						UUID: groupUUIDs[0],
					},
					{
						UUID:            groupUUIDs[1],
						HasActiveEvents: true,
					},
					{
						UUID:            groupUUIDs[2],
						HasActiveEvents: true,
					},
				}, nil),
				traceRepo: nil,
				eventRepo: nil,
			},
			args: args{
				ctx: context.Background(),
			},
			want: []*models.GroupPreview{
				{
					UUID:            groupUUIDs[1],
					HasActiveEvents: true,
				},
				{
					UUID:            groupUUIDs[2],
					HasActiveEvents: true,
				},
			},
			want1: []*models.GroupPreview{
				{
					UUID: groupUUIDs[0],
				},
			},
			wantErr: false,
		},
		{
			name: "error get group list",
			fields: fields{
				groupRepo: mocks.NewGroupRepositoryMock(mc).GetListMock.Return(nil, errors.New("error")),
				traceRepo: nil,
				eventRepo: nil,
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ShowManager{
				groupRepo: tt.fields.groupRepo,
				traceRepo: tt.fields.traceRepo,
				eventRepo: tt.fields.eventRepo,
			}
			got, got1, err := m.GetGroupList(tt.args.ctx, models.Filters{})
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGroupList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !equalGroupLists(got, tt.want) {
				t.Errorf("GetGroupList() got = %v, want %v", got, tt.want)
			}
			if !equalGroupLists(got1, tt.want1) {
				t.Errorf("GetGroupList() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func equalTraces(t1, t2 models.Trace) bool {
	if t1.UUID.String() != t2.UUID.String() {
		return false
	}

	if len(t1.Children) != len(t2.Children) {
		return false
	}

	if len(t1.Events) != len(t2.Events) {
		return false
	}

	for k := range t1.Events {
		if t1.Events[k].UUID.String() != t2.Events[k].UUID.String() {
			return false
		}
	}

	for i := range t1.Children {
		if !equalTraces(t1.Children[i], t2.Children[i]) {
			return false
		}
	}

	return true
}

func equalGroupLists(l1, l2 []*models.GroupPreview) bool {
	if len(l1) != len(l2) {
		return false
	}

	for i := range l1 {
		g1, g2 := *(l1[i]), *(l2[i])

		if g1.UUID != g2.UUID {
			return false
		}
	}

	return true
}

func TestShowManager_GetGroup(t *testing.T) {
	mc := minimock.NewController(t)

	groupUUID := uuid.New()

	type fields struct {
		groupRepo ports.GroupRepository
		traceRepo ports.TraceRepository
		eventRepo ports.EventRepository
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Group
		wantErr bool
	}{
		{
			name: "success get group",
			fields: fields{
				groupRepo: mocks.NewGroupRepositoryMock(mc).GetMock.Return(&models.Group{
					UUID:   groupUUID,
					Traces: nil,
				}, nil),
				traceRepo: nil,
				eventRepo: nil,
			},
			args: args{
				ctx: context.Background(),
				id:  groupUUID,
			},
			want: &models.Group{
				UUID:   groupUUID,
				Traces: nil,
			},
			wantErr: false,
		},
		{
			name: "error get group",
			fields: fields{
				groupRepo: mocks.NewGroupRepositoryMock(mc).GetMock.Return(nil, errors.New("error")),
				traceRepo: nil,
				eventRepo: nil,
			},
			args: args{
				ctx: context.Background(),
				id:  groupUUID,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ShowManager{
				groupRepo: tt.fields.groupRepo,
				traceRepo: tt.fields.traceRepo,
				eventRepo: tt.fields.eventRepo,
			}
			got, err := m.GetGroup(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGroup() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShowManager_GetEventList(t *testing.T) {
	mc := minimock.NewController(t)

	traceUUID := uuid.New()
	eventUUIDs := []uuid.UUID{
		uuid.New(),
		uuid.New(),
		uuid.New(),
	}

	type fields struct {
		groupRepo ports.GroupRepository
		traceRepo ports.TraceRepository
		eventRepo ports.EventRepository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.Event
		wantErr bool
	}{
		{
			name: "success get events",
			fields: fields{
				groupRepo: nil,
				traceRepo: nil,
				eventRepo: mocks.NewEventRepositoryMock(mc).GetListMock.Return([]*models.Event{
					{
						UUID:      eventUUIDs[0],
						TraceUUID: traceUUID,
						Message:   "message 0",
						Priority:  models.PriorityInfo,
						Fixed:     false,
					},
					{
						UUID:      eventUUIDs[1],
						TraceUUID: traceUUID,
						Message:   "message 1",
						Priority:  models.PriorityFatal,
						Fixed:     false,
					},
					{
						UUID:      eventUUIDs[2],
						TraceUUID: traceUUID,
						Message:   "message 2",
						Priority:  models.PriorityFatal,
						Fixed:     false,
					},
				}, nil),
			},
			args: args{ctx: context.Background()},
			want: []*models.Event{
				{
					UUID:      eventUUIDs[0],
					TraceUUID: traceUUID,
					Message:   "message 0",
					Priority:  models.PriorityInfo,
					Fixed:     false,
				},
				{
					UUID:      eventUUIDs[1],
					TraceUUID: traceUUID,
					Message:   "message 1",
					Priority:  models.PriorityFatal,
					Fixed:     false,
				},
				{
					UUID:      eventUUIDs[2],
					TraceUUID: traceUUID,
					Message:   "message 2",
					Priority:  models.PriorityFatal,
					Fixed:     false,
				},
			},
			wantErr: false,
		},
		{
			name: "error get events",
			fields: fields{
				groupRepo: nil,
				traceRepo: nil,
				eventRepo: mocks.NewEventRepositoryMock(mc).GetListMock.Return(nil, errors.New("error")),
			},
			args:    args{ctx: context.Background()},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ShowManager{
				groupRepo: tt.fields.groupRepo,
				traceRepo: tt.fields.traceRepo,
				eventRepo: tt.fields.eventRepo,
			}
			got, err := m.GetEventList(tt.args.ctx, false)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEventList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEventList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShowManager_GetEvent(t *testing.T) {
	mc := minimock.NewController(t)

	traceUUID := uuid.New()
	eventUUID := uuid.New()

	type fields struct {
		groupRepo ports.GroupRepository
		traceRepo ports.TraceRepository
		eventRepo ports.EventRepository
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Event
		wantErr bool
	}{
		{
			name: "success get event",
			fields: fields{
				groupRepo: nil,
				traceRepo: nil,
				eventRepo: mocks.NewEventRepositoryMock(mc).GetMock.Return(&models.Event{
					UUID:      eventUUID,
					TraceUUID: traceUUID,
					Message:   "aaa",
					Priority:  models.PriorityError,
				}, nil),
			},
			args: args{
				ctx: context.Background(),
				id:  eventUUID,
			},
			want: &models.Event{
				UUID:      eventUUID,
				TraceUUID: traceUUID,
				Message:   "aaa",
				Priority:  models.PriorityError,
			},
			wantErr: false,
		},
		{
			name: "error get event",
			fields: fields{
				groupRepo: nil,
				traceRepo: nil,
				eventRepo: mocks.NewEventRepositoryMock(mc).GetMock.Return(nil, errors.New("error")),
			},
			args: args{
				ctx: context.Background(),
				id:  eventUUID,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ShowManager{
				groupRepo: tt.fields.groupRepo,
				traceRepo: tt.fields.traceRepo,
				eventRepo: tt.fields.eventRepo,
			}
			got, err := m.GetEvent(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEvent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShowManager_FixEvent(t *testing.T) {
	mc := minimock.NewController(t)

	traceUUID := uuid.New()
	eventUUID := uuid.New()

	type fields struct {
		groupRepo ports.GroupRepository
		traceRepo ports.TraceRepository
		eventRepo ports.EventRepository
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success fix event",
			fields: fields{
				groupRepo: nil,
				traceRepo: nil,
				eventRepo: mocks.NewEventRepositoryMock(mc).GetMock.Return(&models.Event{
					UUID:      eventUUID,
					TraceUUID: traceUUID,
					Message:   "message",
					Priority:  models.PriorityFatal,
					Payload:   nil,
					Fixed:     false,
				}, nil).UpdateMock.Return(nil),
			},
			args: args{
				ctx: context.Background(),
				id:  eventUUID,
			},
			wantErr: false,
		},
		{
			name: "error get event",
			fields: fields{
				groupRepo: nil,
				traceRepo: nil,
				eventRepo: mocks.NewEventRepositoryMock(mc).GetMock.Return(nil, errors.New("error")),
			},
			args: args{
				ctx: context.Background(),
				id:  eventUUID,
			},
			wantErr: true,
		},
		{
			name: "error update event",
			fields: fields{
				groupRepo: nil,
				traceRepo: nil,
				eventRepo: mocks.NewEventRepositoryMock(mc).GetMock.Return(&models.Event{
					UUID:      eventUUID,
					TraceUUID: traceUUID,
					Message:   "message",
					Priority:  models.PriorityFatal,
					Payload:   nil,
					Fixed:     false,
				}, nil).UpdateMock.Return(errors.New("error")),
			},
			args: args{
				ctx: context.Background(),
				id:  eventUUID,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ShowManager{
				groupRepo: tt.fields.groupRepo,
				traceRepo: tt.fields.traceRepo,
				eventRepo: tt.fields.eventRepo,
			}
			if err := m.FixEvent(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("FixEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
