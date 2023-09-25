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

type GroupRepository struct {
	cache *redis.Client
	db    ports.GroupRepository
	ttl   time.Duration
}

func NewGroupRepository(cache *redis.Client, db ports.GroupRepository, ttl time.Duration) *GroupRepository {
	return &GroupRepository{
		cache: cache,
		db:    db,
		ttl:   ttl,
	}
}

func (g *GroupRepository) Create(ctx context.Context) (*models.Group, error) {
	group, err := g.db.Create(ctx)
	if err != nil {
		return nil, fmt.Errorf("create group via db: %w", err)
	}

	data, err := json.Marshal(group)
	if err != nil {
		return nil, fmt.Errorf("marshal group: %w", err)
	}

	if err := g.cache.Set(ctx, "group_"+group.UUID.String(), data, g.ttl).Err(); err != nil {
		return nil, fmt.Errorf("cant set to cache: %w", err)
	}

	return group, nil
}

func (g *GroupRepository) Get(ctx context.Context, id uuid.UUID) (*models.Group, error) {
	key := "group_" + id.String()

	if g.cache.Exists(ctx, key).Val() == 0 {
		return g.db.Get(ctx, id)
	}

	data, err := g.cache.Get(ctx, key).Bytes()
	if err != nil {
		return nil, fmt.Errorf("get value from cache: %w", err)
	}

	var group models.Group

	if err := json.Unmarshal(data, &group); err != nil {
		return nil, fmt.Errorf("unmarshal group: %w", err)
	}

	return &group, nil
}

func (g *GroupRepository) GetList(ctx context.Context, filters models.Filters) ([]*models.GroupPreview, error) {
	groups, err := g.db.GetList(ctx, filters)
	if err != nil {
		return nil, fmt.Errorf("get list from db: %w", err)
	}

	return groups, nil
}
