package clickhouse

import (
	"context"
	"fmt"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/jsonenc"
	"github.com/Shopify/sarama"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GroupRepository struct {
	db       *gorm.DB
	producer sarama.SyncProducer
}

func NewGroupRepository(db *gorm.DB, producer sarama.SyncProducer) *GroupRepository {
	return &GroupRepository{db: db, producer: producer}
}

func (g *GroupRepository) Create(ctx context.Context) (*models.Group, error) {
	group := Group{UUID: uuid.New()}

	_, _, err := g.producer.SendMessage(&sarama.ProducerMessage{
		Topic: "groups",
		Key:   sarama.StringEncoder(group.UUID.String()),
		Value: jsonenc.New(group),
	})

	if err != nil {
		return nil, fmt.Errorf("send message: %w", err)
	}

	return &models.Group{
		UUID:   group.UUID,
		Traces: nil,
	}, nil
}

func (g *GroupRepository) Get(ctx context.Context, id uuid.UUID) (*models.Group, error) {
	var group Group

	if err := g.db.WithContext(ctx).Table("group").Where("uuid = ?", id).Take(&group).Error; err != nil {
		return nil, fmt.Errorf("get group from clickhouse: %w", err)
	}

	result, err := toGroup(ctx, g.db, group)
	if err != nil {
		return nil, fmt.Errorf("convert group: %w", err)
	}

	return result, nil
}

func (g *GroupRepository) queryWithoutFilters() *gorm.DB {
	query := `
		select
			gt.group_uuid as uuid,
			gt.start as start,
			gt.end as end,
			gn.title as title,
			not(gf.not_fixed) as fixed
		from
		
		(select
			t.group_uuid as group_uuid,
			min(t.time_start) as start,
			min(coalesce(t.time_end, now())) as end
		from trace t group by t.group_uuid) gt
		
		inner join
		
		(select
			t.group_uuid as group_uuid,
			t.time_start as start,
			t.title as title
		from trace t where parent_uuid = '00000000-0000-0000-0000-000000000000') gn
		
		on gt.start = gn.start and gt.group_uuid = gn.group_uuid
		
		left join
		
		(select
			e.group_uuid as group_uuid,
			not(min(e.fixed)) as not_fixed
		from event e group by group_uuid) gf
		
		on gt.group_uuid = gf.group_uuid
		
		order by start desc
	`

	return g.db.Raw(query)
}

func (g *GroupRepository) queryWithLimitFilter(limit int) *gorm.DB {
	query := `
		select
			gt.group_uuid as uuid,
			gt.start as start,
			gt.end as end,
			gn.title as title,
			not(gf.not_fixed) as fixed
		from
		
		(select
			t.group_uuid as group_uuid,
			min(t.time_start) as start,
			min(coalesce(t.time_end, now())) as end
		from trace t group by t.group_uuid) gt
		
		inner join
		
		(select
			t.group_uuid as group_uuid,
			t.time_start as start,
			t.title as title
		from trace t where parent_uuid = '00000000-0000-0000-0000-000000000000') gn
		
		on gt.start = gn.start and gt.group_uuid = gn.group_uuid
		
		left join
		
		(select
			e.group_uuid as group_uuid,
			not(min(e.fixed)) as not_fixed
		from event e group by group_uuid) gf
		
		on gt.group_uuid = gf.group_uuid
		
		order by start desc
		limit ?;
	`

	return g.db.Raw(query, limit)
}

func (g *GroupRepository) queryWithComponentFilter(component string) *gorm.DB {
	query := `
		select
			gt.group_uuid as uuid,
			gt.start as start,
			gt.end as end,
			gn.title as title,
			not(gf.not_fixed) as fixed
		from
		
		(select
			t.group_uuid as group_uuid,
			min(t.time_start) as start,
			min(coalesce(t.time_end, now())) as end
		from trace t group by t.group_uuid) gt
		
		inner join
		
		(select
			t.group_uuid as group_uuid,
			t.time_start as start,
			t.title as title
		from trace t where component = ? and parent_uuid = '00000000-0000-0000-0000-000000000000') gn
		
		on gt.start = gn.start and gt.group_uuid = gn.group_uuid
		
		left join
		
		(select
			e.group_uuid as group_uuid,
			not(min(e.fixed)) as not_fixed
		from event e group by group_uuid) gf
		
		on gt.group_uuid = gf.group_uuid
		
		order by start desc
	`

	return g.db.Raw(query, component)
}

func (g *GroupRepository) queryWithAllFilter(component string, limit int) *gorm.DB {
	query := `
		select
			gt.group_uuid as uuid,
			gt.start as start,
			gt.end as end,
			gn.title as title,
			not(gf.not_fixed) as fixed
		from
		
		(select
			t.group_uuid as group_uuid,
			min(t.time_start) as start,
			min(coalesce(t.time_end, now())) as end
		from trace t group by t.group_uuid) gt
		
		inner join
		
		(select
			t.group_uuid as group_uuid,
			t.time_start as start,
			t.title as title
		from trace t where component = ? and parent_uuid = '00000000-0000-0000-0000-000000000000') gn
		
		on gt.start = gn.start and gt.group_uuid = gn.group_uuid
		
		left join
		
		(select
			e.group_uuid as group_uuid,
			not(min(e.fixed)) as not_fixed
		from event e group by group_uuid) gf
		
		on gt.group_uuid = gf.group_uuid
		
		order by start desc
		limit ?;
	`

	return g.db.Raw(query, component, limit)
}

func (g *GroupRepository) getQuery(f models.Filters) *gorm.DB {
	if f.Component == "" && f.Limit == 0 {
		return g.queryWithoutFilters()
	}

	if f.Component == "" && f.Limit != 0 {
		return g.queryWithLimitFilter(f.Limit)
	}

	if f.Component != "" && f.Limit == 0 {
		return g.queryWithComponentFilter(f.Component)
	}

	return g.queryWithAllFilter(f.Component, f.Limit)
}

func (g *GroupRepository) GetList(ctx context.Context, filters models.Filters) ([]*models.GroupPreview, error) {
	var groups []GroupPreview

	if err := g.getQuery(filters).WithContext(ctx).Find(&groups).Error; err != nil {
		return nil, fmt.Errorf("get groups from clickhouse: %w", err)
	}

	result := make([]*models.GroupPreview, len(groups))

	err := g.db.Transaction(func(tx *gorm.DB) error {
		for i, group := range groups {
			var err error
			if result[i], err = toGroupPreview(ctx, group, tx); err != nil {
				return fmt.Errorf("convert to group: %w", err)
			}
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("tx: %w", err)
	}

	return result, nil
}
