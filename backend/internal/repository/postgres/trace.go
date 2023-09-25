package postgres

import (
	"context"
	"fmt"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TraceRepository struct {
	db *gorm.DB
}

func NewTraceRepository(db *gorm.DB) *TraceRepository {
	return &TraceRepository{db: db}
}

func toStringPtr(s string) *string {
	return &s
}

func (t *TraceRepository) Create(ctx context.Context, trace *models.Trace) (*models.Trace, error) {
	tr := Trace{
		UUID:       uuid.New(),
		GroupUUID:  trace.GroupUUID,
		ParentUUID: trace.ParentUUID,
		Title:      trace.Title,
		TimeStart:  trace.Start,
		Component:  trace.Component,
	}

	if err := t.db.WithContext(ctx).Table("trace").Create(&tr).Error; err != nil {
		return nil, fmt.Errorf("create trace: %w", err)
	}

	return &models.Trace{
		UUID:       tr.UUID,
		GroupUUID:  tr.GroupUUID,
		ParentUUID: tr.ParentUUID,
		Start:      tr.TimeStart,
		Title:      tr.Title,
		Component:  tr.Component,
	}, nil
}

func (t *TraceRepository) Update(ctx context.Context, trace *models.Trace) error {
	var emptyTimeEnd bool
	if err := t.db.WithContext(ctx).Table("trace").Select("time_end is null").Where("uuid = ?", trace.UUID).Scan(&emptyTimeEnd).Error; err != nil {
		return fmt.Errorf("get end time from repo: %w", err)
	}

	tr := Trace{
		UUID:       trace.UUID,
		GroupUUID:  trace.GroupUUID,
		ParentUUID: trace.ParentUUID,
		Title:      trace.Title,
		TimeStart:  trace.Start,
		TimeEnd:    &trace.End,
		Component:  trace.Component,
	}

	if err := t.db.WithContext(ctx).Table("trace").Save(&tr).Error; err != nil {
		return fmt.Errorf("save trace in db: %w", err)
	}

	return nil
}

func (t *TraceRepository) Get(ctx context.Context, id uuid.UUID) (*models.Trace, error) {
	var tr Trace
	if err := t.db.WithContext(ctx).Table("trace").Where("uuid = ?", id).Take(&tr).Error; err != nil {
		return nil, fmt.Errorf("get trace from db: %w", err)
	}

	trace, err := toTrace(ctx, t.db, tr)
	if err != nil {
		return nil, fmt.Errorf("convert trace: %w", err)
	}

	return &trace, nil
}
