package redisdb

import (
	"context"

	"github.com/go-redis/redis/v8"

	"github.com/IDarar/hub/internal/config"
	"github.com/IDarar/hub/pkg/logger"
)

var ctx = context.Background()

func NewRedisDB(cfg *config.Config) (*redis.Client, error) {

	rdb := redis.NewClient(&redis.Options{
		//TODO config
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	status := rdb.Ping(ctx)
	if status.Err() != nil {
		logger.Error("not connected to redis", status.Err())
		return nil, status.Err()
	}
	return rdb, status.Err()

}
