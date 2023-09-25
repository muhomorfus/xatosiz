package managers

import (
	"context"
	"errors"
	"testing"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/ports"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/repository/postgres"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/testdb"
	"github.com/google/uuid"
)

func checkContainsEvent(message string, priority models.Priority) func(m *AcceptManager) error {
	return func(m *AcceptManager) error {
		list, err := m.eventRepo.GetList(context.Background(), false)
		if err != nil {
			return err
		}

		for _, e := range list {
			if e.Message == message && e.Priority == priority {
				return nil
			}
		}

		return errors.New("event not in event list")
	}
}

func TestIntegrationAcceptManager_CreateEvent(t *testing.T) {
	c, db, err := testdb.New()
	if err != nil {
		t.Fatal(err)
	}
	defer c.Terminate(context.Background())

	gr := postgres.NewGroupRepository(db)
	tr := postgres.NewTraceRepository(db)
	er := postgres.NewEventRepository(db)

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
		check   func(m *AcceptManager) error
	}{
		{
			name: "success empty trace id",
			fields: fields{
				groupRepo: gr,
				traceRepo: tr,
				eventRepo: er,
			},
			args: args{
				ctx: context.Background(),
				event: models.EventBaseInfo{
					GroupUUID: toPtr(uuid.MustParse("8e5ba9ce-1d63-4bfb-92f4-15fa1b7f71b5")),
					TraceUUID: nil,
					Message:   "Some event",
					Priority:  models.PriorityInfo,
					Payload:   nil,
					Component: "component",
				},
			},
			want: &models.Event{
				Message:  "Some event",
				Priority: models.PriorityInfo,
			},
			wantErr: false,
			check:   checkContainsEvent("Some event", models.PriorityInfo),
		},
		{
			name: "success empty group id",
			fields: fields{
				groupRepo: gr,
				traceRepo: tr,
				eventRepo: er,
			},
			args: args{
				ctx: context.Background(),
				event: models.EventBaseInfo{
					GroupUUID: nil,
					TraceUUID: nil,
					Message:   "Cool event",
					Priority:  models.PriorityError,
					Payload:   nil,
					Component: "component",
				},
			},
			want: &models.Event{
				Message:  "Cool event",
				Priority: models.PriorityError,
			},
			wantErr: false,
			check:   checkContainsEvent("Cool event", models.PriorityError),
		},
		{
			name: "success exist trace",
			fields: fields{
				groupRepo: gr,
				traceRepo: tr,
				eventRepo: er,
			},
			args: args{
				ctx: context.Background(),
				event: models.EventBaseInfo{
					GroupUUID: nil,
					TraceUUID: toPtr(uuid.MustParse("888eaedd-7855-4a9c-b6ab-45ba76123b01")),
					Message:   "Super event",
					Priority:  models.PriorityFatal,
					Payload:   nil,
					Component: "component",
				},
			},
			want: &models.Event{
				Message:  "Super event",
				Priority: models.PriorityFatal,
			},
			wantErr: false,
			check:   checkContainsEvent("Super event", models.PriorityFatal),
		},
		{
			name: "trace not exist",
			fields: fields{
				groupRepo: gr,
				traceRepo: tr,
				eventRepo: er,
			},
			args: args{
				ctx: context.Background(),
				event: models.EventBaseInfo{
					GroupUUID: nil,
					TraceUUID: toPtr(uuid.MustParse("888eaedd-7855-4a9c-b6ab-45ba76123b02")),
					Message:   "Super event",
					Priority:  models.PriorityFatal,
					Payload:   nil,
					Component: "component",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "trace is finished",
			fields: fields{
				groupRepo: gr,
				traceRepo: tr,
				eventRepo: er,
			},
			args: args{
				ctx: context.Background(),
				event: models.EventBaseInfo{
					GroupUUID: nil,
					TraceUUID: toPtr(uuid.MustParse("9be1f78f-6df3-43ea-ab0c-f7c58c6e675b")),
					Message:   "Super event",
					Priority:  models.PriorityFatal,
					Payload:   nil,
					Component: "component",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "group not exist",
			fields: fields{
				groupRepo: gr,
				traceRepo: tr,
				eventRepo: er,
			},
			args: args{
				ctx: context.Background(),
				event: models.EventBaseInfo{
					GroupUUID: toPtr(uuid.MustParse("9be1f78f-6df3-43ea-ab0c-f7c58c6e675b")),
					TraceUUID: nil,
					Message:   "Super event",
					Priority:  models.PriorityFatal,
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
			if (err == nil) && (got.Message != tt.want.Message || got.Priority != tt.want.Priority) {
				t.Errorf("CreateEvent() got = %v, want %v", got, tt.want)
			}
			if tt.check != nil {
				if err := tt.check(m); err != nil {
					t.Errorf("got unexpected check error = %v", err)
				}
			}
		})
	}
}
