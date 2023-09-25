package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/ports"
	"github.com/redis/go-redis/v9"
	"time"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"github.com/google/uuid"
)

type TraceRepository struct {
	cache *redis.Client
	db    ports.TraceRepository
	ttl   time.Duration
}

func NewTraceRepository(cache *redis.Client, db ports.TraceRepository, ttl time.Duration) *TraceRepository {
	return &TraceRepository{
		cache: cache,
		db:    db,
		ttl:   ttl,
	}
}

func (t *TraceRepository) Create(ctx context.Context, trace *models.Trace) (*models.Trace, error) {
	trace, err := t.db.Create(ctx, trace)
	if err != nil {
		return nil, fmt.Errorf("create trace via db: %w", err)
	}

	data, err := json.Marshal(trace)
	if err != nil {
		return nil, fmt.Errorf("marshal trace: %w", err)
	}

	if err := t.cache.Set(ctx, "trace_"+trace.UUID.String(), data, t.ttl).Err(); err != nil {
		return nil, fmt.Errorf("cant set to cache: %w", err)
	}

	return trace, nil
}

func (t *TraceRepository) Update(ctx context.Context, trace *models.Trace) error {
	key := "trace_" + trace.UUID.String()

	if t.cache.Exists(ctx, key).Val() == 0 {
		return t.db.Update(ctx, trace)
	}

	data, err := t.cache.Get(ctx, key).Bytes()
	if err != nil {
		return fmt.Errorf("get trace from cache: %w", err)
	}

	var traceStored models.Trace

	if err := json.Unmarshal(data, &traceStored); err != nil {
		return fmt.Errorf("unmarshal trace: %w", err)
	}

	data, err = json.Marshal(trace)
	if err != nil {
		return fmt.Errorf("marshal trace: %w", err)
	}

	if err := t.cache.Set(ctx, key, data, t.ttl).Err(); err != nil {
		return fmt.Errorf("cant set to cache: %w", err)
	}

	if err := t.db.Update(ctx, trace); err != nil {
		return fmt.Errorf("update trace in db: %w", err)
	}

	return nil
}

func (t *TraceRepository) Get(ctx context.Context, id uuid.UUID) (*models.Trace, error) {
	key := "trace_" + id.String()

	if t.cache.Exists(ctx, key).Val() == 0 {
		return t.db.Get(ctx, id)
	}

	data, err := t.cache.Get(ctx, key).Bytes()
	if err != nil {
		return nil, fmt.Errorf("get trace from cache: %w", err)
	}

	var trace models.Trace

	if err := json.Unmarshal(data, &trace); err != nil {
		return nil, fmt.Errorf("unmarshal trace: %w", err)
	}

	return &trace, nil
}
