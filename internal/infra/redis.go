package infra

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
	"haejoong.com-api/internal/config"
)

func NewRedisClient(redisConfig config.RedisConfig) (*redis.Client, error) {
	addr := fmt.Sprintf(
		"%s:%s",
		redisConfig.Host,
		redisConfig.Port,
	)

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: redisConfig.Pass,
		DB:       redisConfig.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		slog.Error("Redis 연결 실패",
			"error", err,
			"addr", addr,
		)
		return nil, err
	}

	return rdb, nil
}
