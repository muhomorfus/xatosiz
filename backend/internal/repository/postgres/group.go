package postgres

import (
	"context"
	"fmt"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GroupRepository struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) *GroupRepository {
	return &GroupRepository{db: db}
}

func (g *GroupRepository) Create(ctx context.Context) (*models.Group, error) {
	group := Group{UUID: uuid.New()}
	if err := g.db.WithContext(ctx).Table("group").Create(&group).Error; err != nil {
		return nil, fmt.Errorf("create group via postgres: %w", err)
	}

	return &models.Group{
		UUID:   group.UUID,
		Traces: nil,
	}, nil
}

func (g *GroupRepository) Get(ctx context.Context, id uuid.UUID) (*models.Group, error) {
	var group Group

	if err := g.db.WithContext(ctx).Table("group").Where("uuid = ?", id).Take(&group).Error; err != nil {
		return nil, fmt.Errorf("get group from postgres: %w", err)
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
			g_agr.uuid as uuid,
			g_agr.start as start,
			g_agr."end" as "end",
			(select count(*) = 0 from event e where (select group_uuid from trace where uuid = e.trace_uuid) = g_agr.uuid and e.fixed = false) as fixed,
			(select title from trace where time_start = g_agr.start limit 1) as title
		from (
			 select uuid,
					(select min(t.time_start) from trace t where group_uuid = g.uuid) as start,
					(select max(t.time_end) from (select tt.group_uuid as group_uuid, coalesce(tt.time_end, current_timestamp) as time_end from trace tt where tt.group_uuid = g.uuid) t)::timestamp as "end"
			 from "group" g
			 where (select count(*) from trace where group_uuid = g.uuid) != 0
		) g_agr
		order by g_agr."start" desc;
	`

	return g.db.Raw(query)
}

func (g *GroupRepository) queryWithLimitFilter(limit int) *gorm.DB {
	query := `
		select
			g_agr.uuid as uuid,
			g_agr.start as start,
			g_agr."end" as "end",
			(select count(*) = 0 from event e where (select group_uuid from trace where uuid = e.trace_uuid) = g_agr.uuid and e.fixed = false) as fixed,
			(select title from trace where time_start = g_agr.start limit 1) as title
		from (
			 select uuid,
					(select min(t.time_start) from trace t where group_uuid = g.uuid) as start,
					(select max(t.time_end) from (select tt.group_uuid as group_uuid, coalesce(tt.time_end, current_timestamp) as time_end from trace tt where tt.group_uuid = g.uuid) t)::timestamp as "end"
			 from "group" g
			 where (select count(*) from trace where group_uuid = g.uuid) != 0
		) g_agr
		order by g_agr."start" desc
		limit ?;
	`

	return g.db.Raw(query, limit)
}

func (g *GroupRepository) queryWithComponentFilter(component string) *gorm.DB {
	query := `
		select
			g_agr.uuid as uuid,
			g_agr.start as start,
			g_agr."end" as "end",
			(select count(*) = 0 from event e where (select group_uuid from trace where uuid = e.trace_uuid) = g_agr.uuid and e.fixed = false) as fixed,
			(select title from trace where time_start = g_agr.start limit 1) as title
		from (
			 select uuid,
					(select min(t.time_start) from trace t where group_uuid = g.uuid) as start,
					(select max(t.time_end) from (select tt.group_uuid as group_uuid, coalesce(tt.time_end, current_timestamp) as time_end from trace tt where tt.group_uuid = g.uuid) t)::timestamp as "end"
			 from "group" g
			 where (select count(*) from trace where group_uuid = g.uuid and component = ?) != 0
		) g_agr
		order by g_agr."start" desc;
	`

	return g.db.Raw(query, component)
}

func (g *GroupRepository) queryWithAllFilter(component string, limit int) *gorm.DB {
	query := `
		select
			g_agr.uuid as uuid,
			g_agr.start as start,
			g_agr."end" as "end",
			(select count(*) = 0 from event e where (select group_uuid from trace where uuid = e.trace_uuid) = g_agr.uuid and e.fixed = false) as fixed,
			(select title from trace where time_start = g_agr.start limit 1) as title
		from (
			 select uuid,
					(select min(t.time_start) from trace t where group_uuid = g.uuid) as start,
					(select max(t.time_end) from (select tt.group_uuid as group_uuid, coalesce(tt.time_end, current_timestamp) as time_end from trace tt where tt.group_uuid = g.uuid) t)::timestamp as "end"
			 from "group" g
			 where (select count(*) from trace where group_uuid = g.uuid and component = ?) != 0
		) g_agr
		order by g_agr."start" desc
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
		return nil, fmt.Errorf("get groups from postgres: %w", err)
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
