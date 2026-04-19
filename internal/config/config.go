package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type DBConfig struct {
	Host string `env:"DB_HOST" env-default:"localhost"`
	Port string `env:"DB_PORT" env-default:"5432"`
	User string `env:"DB_USER" env-required:"true"`
	Pass string `env:"DB_PASS" env-required:"true"`
	Name string `env:"DB_NAME" env-required:"true"`
}

type RedisConfig struct {
	Host string `env:"REDIS_HOST" env-required:"true"`
	Port string `env:"REDIS_PORT" env-default:"6379"`
	Pass string `env:"REDIS_PASS" env-default:""`
	DB   int    `env:"REDIS_DB" env-default:"0"`
}

type LogConfig struct {
	LogLevel    string `env:"LOG_LEVEL" env-default:"INFO"`
	LogFilePath string `env:"LOG_FILE_PATH" env-default:"./storage/logs/api.log"`
}

type Config struct {
	DB    DBConfig
	Redis RedisConfig
	Log   LogConfig
}

func Load() (*Config, error) {
	cfg := &Config{}

	_ = cleanenv.ReadConfig(".env", cfg)
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
