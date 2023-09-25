package managers

import (
	"context"
	"testing"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/ports"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/repository/postgres"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/testdb"
	"github.com/google/uuid"
)

func TestIntegrationShowManager_GetGroupList(t *testing.T) {
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
				groupRepo: gr,
				traceRepo: tr,
				eventRepo: er,
			},
			args: args{
				ctx: context.Background(),
			},
			want: []*models.GroupPreview{
				{
					UUID: uuid.MustParse("b4b17b07-ecd4-45ba-84cc-9708be9eba6f"), // 1
				},
				{
					UUID: uuid.MustParse("567af5eb-192a-4dc0-966a-d5755e248307"), // 2
				},
			},
			want1: []*models.GroupPreview{
				{
					UUID: uuid.MustParse("8e5ba9ce-1d63-4bfb-92f4-15fa1b7f71b5"), // 4
				},
				{
					UUID: uuid.MustParse("f9d5aca1-becb-4812-95a4-659e8d1d9160"), // 3
				},
			},
			wantErr: false,
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
				t.Errorf("GetGroupList() got = %+v, want %+v", got, tt.want)
			}
			if !equalGroupLists(got1, tt.want1) {
				t.Errorf("GetGroupList() got1 = %+v, want %+v", got1, tt.want1)
			}
		})
	}
}
