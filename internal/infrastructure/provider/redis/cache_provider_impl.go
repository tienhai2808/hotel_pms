package redis

import (
	"context"
	"time"

	"github.com/InstayPMS/backend/internal/application/port"
	"github.com/redis/go-redis/v9"
)

type cacheProviderImpl struct {
	rdb *redis.Client
}

func NewCacheProvider(rdb *redis.Client) port.CacheProvider {
	return &cacheProviderImpl{rdb}
}

func (p *cacheProviderImpl) SetString(ctx context.Context, key, str string, ttl time.Duration) error {
	return p.rdb.Set(ctx, key, str, ttl).Err()
}

func (p *cacheProviderImpl) GetString(ctx context.Context, key string) (string, error) {
	str, err := p.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}

	return str, nil
}

func (p *cacheProviderImpl) SetObject(ctx context.Context, key string, data []byte, ttl time.Duration) error {
	return p.rdb.Set(ctx, key, data, ttl).Err()
}

func (p *cacheProviderImpl) GetObject(ctx context.Context, key string) ([]byte, error) {
	data, err := p.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return []byte(data), nil
}

func (p *cacheProviderImpl) Del(ctx context.Context, key string) error {
	return p.rdb.Del(ctx, key).Err()
}

func (p *cacheProviderImpl) GetInt(ctx context.Context, key string) (int, error) {
	num, err := p.rdb.Get(ctx, key).Int()
	if err == redis.Nil {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	return num, nil
}
