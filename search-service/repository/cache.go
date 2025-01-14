package repository

import (
	"encoding/json"
	"github.com/GOAT-prod/goatcontext"
	"github.com/redis/go-redis/v9"
	"search-service/domain"
	"time"
)

const _defaultTTL = 5 * time.Minute

type Cache interface {
	Get(ctx goatcontext.Context, searchId string) (products []domain.Product, err error)
	Set(ctx goatcontext.Context, searchId string, products []domain.Product) error
	Check(ctx goatcontext.Context, searchId string) bool
}

type CacheRepository struct {
	redis *redis.Client
}

func NewCacheRepository(redis *redis.Client) Cache {
	return &CacheRepository{
		redis: redis,
	}
}

func (r *CacheRepository) Get(ctx goatcontext.Context, searchId string) (products []domain.Product, err error) {
	dataBytes, err := r.redis.Get(ctx, searchId).Bytes()
	if err != nil {
		return nil, err
	}

	return products, json.Unmarshal(dataBytes, &products)
}

func (r *CacheRepository) Set(ctx goatcontext.Context, searchId string, products []domain.Product) error {
	dataBytes, err := json.Marshal(products)
	if err != nil {
		return err
	}

	return r.redis.Set(ctx, searchId, dataBytes, _defaultTTL).Err()
}

func (r *CacheRepository) Check(ctx goatcontext.Context, searchId string) bool {
	value := r.redis.Exists(ctx, searchId)
	exists, err := value.Result()
	if err != nil {
		return false
	}

	return exists == 1
}
