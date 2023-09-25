package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"sort"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"github.com/google/uuid"
)

type AlertRepository struct {
	cache *redis.Client
}

func NewAlertRepository(cache *redis.Client) *AlertRepository {
	return &AlertRepository{
		cache: cache,
	}
}

func keyPrefix(solved bool) string {
	if solved {
		return "alert_solved_"
	}

	return "alert_unsolved_"
}

func key(id uuid.UUID, solved bool) string {
	return keyPrefix(solved) + id.String()
}

func (a *AlertRepository) Create(ctx context.Context, alert *models.Alert) error {
	data, err := json.Marshal(alert)
	if err != nil {
		return fmt.Errorf("marshal alert: %w", err)
	}

	if err := a.cache.Set(ctx, key(alert.UUID, false), data, 0).Err(); err != nil {
		return fmt.Errorf("set alert to redis: %w", err)
	}

	return nil
}

func (a *AlertRepository) GetList(ctx context.Context, solved bool) ([]models.Alert, error) {
	keys, err := a.cache.Keys(ctx, keyPrefix(solved)+"*").Result()
	if err != nil {
		return nil, fmt.Errorf("get alert keys: %w", err)
	}

	result := make([]models.Alert, len(keys))
	for i, key := range keys {
		data, err := a.cache.Get(ctx, key).Bytes()
		if err != nil {
			return nil, fmt.Errorf("get alert: %w", err)
		}

		if err := json.Unmarshal(data, &result[i]); err != nil {
			return nil, fmt.Errorf("unmarshal alert")
		}
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].Time.After(result[j].Time)
	})

	return result, nil
}

func (a *AlertRepository) MarkSolved(ctx context.Context, id uuid.UUID) error {
	data, err := a.cache.Get(ctx, key(id, false)).Bytes()
	if err != nil {
		return fmt.Errorf("get alert from redis: %w", err)
	}

	if err := a.cache.Set(ctx, key(id, true), data, 0).Err(); err != nil {
		return fmt.Errorf("set alert to redis: %w", err)
	}

	if err := a.cache.Del(ctx, key(id, false)).Err(); err != nil {
		return fmt.Errorf("delete key from redis: %w", err)
	}

	return nil
}
