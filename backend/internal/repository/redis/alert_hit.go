package redis

import (
	"context"
	"fmt"
	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/ports"
	"github.com/redis/go-redis/v9"
	"time"

	"github.com/google/uuid"
)

type AlertHitRepository struct {
	cache *redis.Client
	db    ports.AlertHitRepository
	ttl   time.Duration
}

func NewAlertHitRepository(cache *redis.Client, db ports.AlertHitRepository, ttl time.Duration) *AlertHitRepository {
	return &AlertHitRepository{
		cache: cache,
		db:    db,
		ttl:   ttl,
	}
}

func (a *AlertHitRepository) Get(ctx context.Context, id uuid.UUID, duration time.Duration) (int, error) {
	hits, err := a.db.Get(ctx, id, duration)
	if err != nil {
		return 0, fmt.Errorf("get db hits: %w", err)
	}

	fromKey := "hit_" + id.String() + "_" + time.Now().UTC().Add(-duration).Format(time.RFC3339Nano)
	cacheHits, err := a.cache.Keys(ctx, id.String()+"*").Result()
	if err != nil {
		return 0, fmt.Errorf("get cache hits: %w", err)
	}

	for _, key := range cacheHits {
		if key > fromKey {
			hits++
		}
	}

	return hits, nil
}

func (a *AlertHitRepository) Create(ctx context.Context, id uuid.UUID) error {
	if err := a.db.Create(ctx, id); err != nil {
		return fmt.Errorf("create hit in db: %w", err)
	}

	if err := a.cache.Set(ctx, "hit_"+id.String()+"_"+time.Now().UTC().Format(time.RFC3339Nano), 1, a.ttl).Err(); err != nil {
		return fmt.Errorf("set value to cache: %w", err)
	}

	return nil
}

func (a *AlertHitRepository) DeleteAll(ctx context.Context, id uuid.UUID) error {
	return a.db.DeleteAll(ctx, id)
}
