package managers

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/mocks"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/ports"
	"github.com/gojuno/minimock/v3"
	"github.com/google/uuid"
)

func TestAcceptManager_StartTrace(t *testing.T) {
	mc := minimock.NewController(t)

	groupUUID := uuid.New()
	traceUUID := uuid.New()
	parentUUID := uuid.New()

	type fields struct {
		groupRepo ports.GroupRepository
		traceRepo ports.TraceRepository
		eventRepo ports.EventRepository
	}
	type args struct {
		ctx   context.Context
		trace models.TraceBaseInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Trace
		wantErr bool
	}{
		{
			name: "success empty group and empty parent id",
			fields: fields{
				groupRepo: mocks.NewGroupRepositoryMock(mc).CreateMock.Return(&models.Group{
					UUID:   groupUUID,
					Traces: nil,
				}, nil),
				traceRepo: mocks.NewTraceRepositoryMock(mc).CreateMock.Return(&models.Trace{
					UUID:       traceUUID,
					GroupUUID:  groupUUID,
					ParentUUID: nil,
					Start:      time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
					End:        time.Time{},
					Title:      "title",
					Component:  "component",
					Children:   nil,
					Events:     nil,
					Finished:   false,
				}, nil),
				eventRepo: nil,
			},
			args: args{
				ctx: context.Background(),
				trace: models.TraceBaseInfo{
					GroupUUID:  nil,
					ParentUUID: nil,
					Title:      "title",
					Component:  "component",
				},
			},
			want: &models.Trace{
				UUID:       traceUUID,
				GroupUUID:  groupUUID,
				ParentUUID: nil,
				Start:      time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
				End:        time.Time{},
				Title:      "title",
				Component:  "component",
				Children:   nil,
				Events:     nil,
				Finished:   false,
			},
			wantErr: false,
		},
		{
			name: "success empty group and not empty parent id",
			fields: fields{
				groupRepo: mocks.NewGroupRepositoryMock(mc).GetMock.Return(&models.Group{}, nil),
				traceRepo: mocks.NewTraceRepositoryMock(mc).GetMock.Return(&models.Trace{
					UUID:       parentUUID,
					GroupUUID:  groupUUID,
					ParentUUID: nil,
					Start:      time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
					End:        time.Time{},
					Title:      "parent title",
					Component:  "component",
					Children:   nil,
					Events:     nil,
					Finished:   false,
				}, nil).CreateMock.Return(&models.Trace{
					UUID:       traceUUID,
					GroupUUID:  groupUUID,
					ParentUUID: &parentUUID,
					Start:      time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
					End:        time.Time{},
					Title:      "title",
					Component:  "component",
					Children:   nil,
					Events:     nil,
					Finished:   false,
				}, nil),
				eventRepo: nil,
			},
			args: args{
				ctx: context.Background(),
				trace: models.TraceBaseInfo{
					GroupUUID:  nil,
					ParentUUID: &parentUUID,
					Title:      "title",
					Component:  "component",
				},
			},
			want: &models.Trace{
				UUID:       traceUUID,
				GroupUUID:  groupUUID,
				ParentUUID: &parentUUID,
				Start:      time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
				End:        time.Time{},
				Title:      "title",
				Component:  "component",
				Children:   nil,
				Events:     nil,
				Finished:   false,
			},
			wantErr: false,
		},
		{
			name: "success not empty group and empty parent id",
			fields: fields{
				groupRepo: mocks.NewGroupRepositoryMock(mc).GetMock.Return(&models.Group{}, nil),
				traceRepo: mocks.NewTraceRepositoryMock(mc).CreateMock.Return(&models.Trace{
					UUID:       traceUUID,
					GroupUUID:  groupUUID,
					ParentUUID: nil,
					Start:      time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
					End:        time.Time{},
					Title:      "title",
					Component:  "component",
					Children:   nil,
					Events:     nil,
					Finished:   false,
				}, nil),
				eventRepo: nil,
			},
			args: args{
				ctx: context.Background(),
				trace: models.TraceBaseInfo{
					GroupUUID:  &groupUUID,
					ParentUUID: nil,
					Title:      "title",
					Component:  "component",
				},
			},
			want: &models.Trace{
				UUID:       traceUUID,
				GroupUUID:  groupUUID,
				ParentUUID: nil,
				Start:      time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
				End:        time.Time{},
				Title:      "title",
				Component:  "component",
				Children:   nil,
				Events:     nil,
				Finished:   false,
			},
			wantErr: false,
		},
		{
			name: "not empty parent with error trace get",
			fields: fields{
				groupRepo: nil,
				traceRepo: mocks.NewTraceRepositoryMock(mc).GetMock.Return(nil, errors.New("error")),
				eventRepo: nil,
			},
			args: args{
				ctx: context.Background(),
				trace: models.TraceBaseInfo{
					GroupUUID:  nil,
					ParentUUID: &parentUUID,
					Title:      "title",
					Component:  "component",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "not empty parent with already finished trace",
			fields: fields{
				groupRepo: nil,
				traceRepo: mocks.NewTraceRepositoryMock(mc).GetMock.Return(&models.Trace{
					UUID:       parentUUID,
					GroupUUID:  groupUUID,
					ParentUUID: nil,
					Start:      time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
					End:        time.Date(2023, 01, 01, 0, 2, 0, 0, time.UTC),
					Title:      "parent title",
					Component:  "component",
					Children:   nil,
					Events:     nil,
					Finished:   true,
				}, nil),
				eventRepo: nil,
			},
			args: args{
				ctx: context.Background(),
				trace: models.TraceBaseInfo{
					GroupUUID:  nil,
					ParentUUID: &parentUUID,
					Title:      "title",
					Component:  "component",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty parent and group error at create group",
			fields: fields{
				groupRepo: mocks.NewGroupRepositoryMock(mc).CreateMock.Return(nil, errors.New("error")),
				traceRepo: nil,
				eventRepo: nil,
			},
			args: args{
				ctx: context.Background(),
				trace: models.TraceBaseInfo{
					GroupUUID:  nil,
					ParentUUID: nil,
					Title:      "title",
					Component:  "component",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "get group error at non-empty group",
			fields: fields{
				groupRepo: mocks.NewGroupRepositoryMock(mc).GetMock.Return(nil, errors.New("error")),
				traceRepo: nil,
				eventRepo: nil,
			},
			args: args{
				ctx: context.Background(),
				trace: models.TraceBaseInfo{
					GroupUUID:  &groupUUID,
					ParentUUID: nil,
					Title:      "title",
					Component:  "component",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "create trace error",
			fields: fields{
				groupRepo: mocks.NewGroupRepositoryMock(mc).GetMock.Return(&models.Group{}, nil),
				traceRepo: mocks.NewTraceRepositoryMock(mc).CreateMock.Return(nil, errors.New("error")),
				eventRepo: nil,
			},
			args: args{
				ctx: context.Background(),
				trace: models.TraceBaseInfo{
					GroupUUID:  &groupUUID,
					ParentUUID: nil,
					Title:      "title",
					Component:  "component",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &AcceptManager{
				groupRepo: tt.fields.groupRepo,
				traceRepo: tt.fields.traceRepo,
				eventRepo: tt.fields.eventRepo,
			}
			got, err := m.StartTrace(tt.args.ctx, tt.args.trace)
			if (err != nil) != tt.wantErr {
				t.Errorf("StartTrace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StartTrace() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAcceptManager_CreateGroup(t *testing.T) {
	mc := minimock.NewController(t)

	groupUUID := uuid.New()

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
		want    *models.Group
		wantErr bool
	}{
		{
			name: "success create group",
			fields: fields{
				groupRepo: mocks.NewGroupRepositoryMock(mc).CreateMock.Return(&models.Group{
					UUID:   groupUUID,
					Traces: nil,
				}, nil),
				traceRepo: nil,
				eventRepo: nil,
			},
			args: args{
				ctx: context.Background(),
			},
			want: &models.Group{
				UUID:   groupUUID,
				Traces: nil,
			},
			wantErr: false,
		},
		{
			name: "error create group",
			fields: fields{
				groupRepo: mocks.NewGroupRepositoryMock(mc).CreateMock.Return(nil, errors.New("error")),
				traceRepo: nil,
				eventRepo: nil,
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &AcceptManager{
				groupRepo: tt.fields.groupRepo,
				traceRepo: tt.fields.traceRepo,
				eventRepo: tt.fields.eventRepo,
			}
			got, err := m.CreateGroup(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateGroup() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAcceptManager_CreateEvent(t *testing.T) {
	mc := minimock.NewController(t)

	groupUUID := uuid.New()
	traceUUID := uuid.New()
	eventUUID := uuid.New()

	type fields struct {
		groupRepo ports.GroupRepository
		traceRepo ports.TraceRepository
		eventRepo ports.EventRepository
	}
	type args struct {
		ctx   context.Context
		event models.EventBaseInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Event
		wantErr bool
	}{
		{
			name: "success empty trace id",
			fields: fields{
				groupRepo: mocks.NewGroupRepositoryMock(mc).GetMock.Return(&models.Group{}, nil),
				traceRepo: mocks.NewTraceRepositoryMock(mc).CreateMock.Return(&models.Trace{
					UUID:       traceUUID,
					GroupUUID:  groupUUID,
					ParentUUID: nil,
					Start:      time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
					End:        time.Time{},
					Title:      "title",
					Component:  "component",
					Children:   nil,
					Events:     nil,
					Finished:   false,
				}, nil),
				eventRepo: mocks.NewEventRepositoryMock(mc).CreateMock.Return(&models.Event{
					UUID:      eventUUID,
					TraceUUID: traceUUID,
					Message:   "hello world",
					Priority:  models.PriorityInfo,
					Payload:   nil,
					Fixed:     false,
				}, nil),
			},
			args: args{
				ctx: context.Background(),
				event: models.EventBaseInfo{
					GroupUUID: &groupUUID,
					TraceUUID: nil,
					Message:   "hello world",
					Priority:  models.PriorityInfo,
					Payload:   nil,
					Component: "component",
				},
			},
			want: &models.Event{
				UUID:      eventUUID,
				TraceUUID: traceUUID,
				Message:   "hello world",
				Priority:  models.PriorityInfo,
				Payload:   nil,
				Fixed:     false,
			},
			wantErr: false,
		},
		{
			name: "empty trace id and error create",
			fields: fields{
				groupRepo: mocks.NewGroupRepositoryMock(mc).GetMock.Return(&models.Group{}, nil),
				traceRepo: mocks.NewTraceRepositoryMock(mc).CreateMock.Return(&models.Trace{
					UUID:       traceUUID,
					GroupUUID:  groupUUID,
					ParentUUID: nil,
					Start:      time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
					End:        time.Time{},
					Title:      "title",
					Component:  "component",
					Children:   nil,
					Events:     nil,
					Finished:   false,
				}, nil),
				eventRepo: mocks.NewEventRepositoryMock(mc).CreateMock.Return(nil, errors.New("error")),
			},
			args: args{
				ctx: context.Background(),
				event: models.EventBaseInfo{
					GroupUUID: &groupUUID,
					TraceUUID: nil,
					Message:   "hello world",
					Priority:  models.PriorityInfo,
					Payload:   nil,
					Component: "component",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error start trace",
			fields: fields{
				groupRepo: mocks.NewGroupRepositoryMock(mc).GetMock.Return(nil, errors.New("error")),
				traceRepo: nil,
				eventRepo: nil,
			},
			args: args{
				ctx: context.Background(),
				event: models.EventBaseInfo{
					GroupUUID: &groupUUID,
					TraceUUID: nil,
					Message:   "hello world",
					Priority:  models.PriorityInfo,
					Payload:   nil,
					Component: "component",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success filled trace id",
			fields: fields{
				groupRepo: mocks.NewGroupRepositoryMock(mc).GetMock.Return(&models.Group{}, nil),
				traceRepo: mocks.NewTraceRepositoryMock(mc).GetMock.Return(&models.Trace{
					UUID:       traceUUID,
					GroupUUID:  groupUUID,
					ParentUUID: nil,
					Start:      time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
					End:        time.Time{},
					Title:      "title",
					Component:  "component",
					Children:   nil,
					Events:     nil,
					Finished:   false,
				}, nil),
				eventRepo: mocks.NewEventRepositoryMock(mc).CreateMock.Return(&models.Event{
					UUID:      eventUUID,
					TraceUUID: traceUUID,
					Message:   "hello world",
					Priority:  models.PriorityInfo,
					Payload:   nil,
					Fixed:     false,
				}, nil),
			},
			args: args{
				ctx: context.Background(),
				event: models.EventBaseInfo{
					GroupUUID: &groupUUID,
					TraceUUID: &traceUUID,
					Message:   "hello world",
					Priority:  models.PriorityInfo,
					Payload:   nil,
					Component: "component",
				},
			},
			want: &models.Event{
				UUID:      eventUUID,
				TraceUUID: traceUUID,
				Message:   "hello world",
				Priority:  models.PriorityInfo,
				Payload:   nil,
				Fixed:     false,
			},
			wantErr: false,
		},
		{
			name: "error get with filled trace id",
			fields: fields{
				groupRepo: mocks.NewGroupRepositoryMock(mc).GetMock.Return(&models.Group{}, nil),
				traceRepo: mocks.NewTraceRepositoryMock(mc).GetMock.Return(nil, errors.New("error")),
				eventRepo: nil,
			},
			args: args{
				ctx: context.Background(),
				event: models.EventBaseInfo{
					GroupUUID: &groupUUID,
					TraceUUID: &traceUUID,
					Message:   "hello world",
					Priority:  models.PriorityInfo,
					Payload:   nil,
					Component: "component",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "finished trace",
			fields: fields{
				groupRepo: mocks.NewGroupRepositoryMock(mc).GetMock.Return(&models.Group{}, nil),
				traceRepo: mocks.NewTraceRepositoryMock(mc).GetMock.Return(&models.Trace{
					UUID:       traceUUID,
					GroupUUID:  groupUUID,
					ParentUUID: nil,
					Start:      time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
					End:        time.Time{},
					Title:      "title",
					Component:  "component",
					Children:   nil,
					Events:     nil,
					Finished:   true,
				}, nil),
				eventRepo: nil,
			},
			args: args{
				ctx: context.Background(),
				event: models.EventBaseInfo{
					GroupUUID: &groupUUID,
					TraceUUID: &traceUUID,
					Message:   "hello world",
					Priority:  models.PriorityInfo,
					Payload:   nil,
					Component: "component",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &AcceptManager{
				groupRepo: tt.fields.groupRepo,
				traceRepo: tt.fields.traceRepo,
				eventRepo: tt.fields.eventRepo,
			}
			got, err := m.CreateEvent(tt.args.ctx, tt.args.event)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateEvent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAcceptManager_EndTrace(t *testing.T) {
	mc := minimock.NewController(t)

	groupUUID := uuid.New()
	traceUUID := uuid.New()

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
			name: "success end trace",
			fields: fields{
				groupRepo: nil,
				traceRepo: mocks.NewTraceRepositoryMock(mc).GetMock.Return(&models.Trace{
					UUID:       traceUUID,
					GroupUUID:  groupUUID,
					ParentUUID: nil,
					Start:      time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
					End:        time.Time{},
					Title:      "title",
					Component:  "component",
					Children:   nil,
					Events:     nil,
					Finished:   false,
				}, nil).UpdateMock.Return(nil),
				eventRepo: nil,
			},
			args: args{
				ctx: context.Background(),
				id:  traceUUID,
			},
			wantErr: false,
		},
		{
			name: "finished trace",
			fields: fields{
				groupRepo: nil,
				traceRepo: mocks.NewTraceRepositoryMock(mc).GetMock.Return(&models.Trace{
					UUID:       traceUUID,
					GroupUUID:  groupUUID,
					ParentUUID: nil,
					Start:      time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
					End:        time.Time{},
					Title:      "title",
					Component:  "component",
					Children:   nil,
					Events:     nil,
					Finished:   true,
				}, nil).UpdateMock.Return(nil),
				eventRepo: nil,
			},
			args: args{
				ctx: context.Background(),
				id:  traceUUID,
			},
			wantErr: true,
		},
		{
			name: "errors with update",
			fields: fields{
				groupRepo: nil,
				traceRepo: mocks.NewTraceRepositoryMock(mc).GetMock.Return(&models.Trace{
					UUID:       traceUUID,
					GroupUUID:  groupUUID,
					ParentUUID: nil,
					Start:      time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC),
					End:        time.Time{},
					Title:      "title",
					Component:  "component",
					Children:   nil,
					Events:     nil,
					Finished:   false,
				}, nil).UpdateMock.Return(errors.New("error")),
				eventRepo: nil,
			},
			args: args{
				ctx: context.Background(),
				id:  traceUUID,
			},
			wantErr: true,
		},
		{
			name: "errors with get",
			fields: fields{
				groupRepo: nil,
				traceRepo: mocks.NewTraceRepositoryMock(mc).GetMock.Return(nil, errors.New("error")),
				eventRepo: nil,
			},
			args: args{
				ctx: context.Background(),
				id:  traceUUID,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &AcceptManager{
				groupRepo: tt.fields.groupRepo,
				traceRepo: tt.fields.traceRepo,
				eventRepo: tt.fields.eventRepo,
			}
			if err := m.EndTrace(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("EndTrace() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
