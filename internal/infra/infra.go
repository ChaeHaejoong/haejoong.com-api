package infra

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"haejoong.com-api/internal/config"
)

type Infra struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func New(cfg *config.Config) (*Infra, error) {
	infra := &Infra{}

	db, err := NewPostgresDB(cfg.DB)
	if err != nil {
		return nil, err
	}

	redisClient, err := NewRedisClient(cfg.Redis)
	if err != nil {
		return nil, err
	}

	infra.DB = db
	infra.Redis = redisClient

	return infra, nil
}
