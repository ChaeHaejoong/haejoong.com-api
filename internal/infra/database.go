package infra

import (
	"fmt"
	"log/slog"

	slogGorm "github.com/orandin/slog-gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"haejoong.com-api/internal/config"
)

func NewPostgresDB(dbConfig config.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Pass,
		dbConfig.Name,
		dbConfig.Port,
	)
	gormLogger := slogGorm.New(slogGorm.WithHandler(slog.Default().Handler()), slogGorm.WithTraceAll())

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: gormLogger})
	if err != nil {
		slog.Error("database 연결 실패",
			"error", err,
			"host", dbConfig.Host,
			"database", dbConfig.Name,
		)
		return nil, err
	}

	return db, nil
}
