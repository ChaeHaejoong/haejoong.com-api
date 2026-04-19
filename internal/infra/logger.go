package infra

import (
	"io"
	"log"
	"log/slog"
	"os"

	"haejoong.com-api/internal/config"
)

// 우선은 환경변수로 로거를 이니셜 (switch case문)
// 로그를 찍는 함수를 만나면 메모리(LogLevel 포인터변수)로 접근해서 현재 값을 파악하고 로그를 찍는 구조
// 그렇기에 런타임에 변수를 바꿀 수 있는 구조

var LogLevel = &slog.LevelVar{}

func InitLogger(cfg config.LogConfig) error {
	envLevel := cfg.LogLevel
	switch envLevel {
	case "DEBUG":
		LogLevel.Set(slog.LevelDebug)
	case "WARN":
		LogLevel.Set(slog.LevelWarn)
	case "ERROR":
		LogLevel.Set(slog.LevelError)
	default:
		LogLevel.Set(slog.LevelInfo)
	}

	file, err := os.OpenFile(cfg.LogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("로그 파일 열기 실패: %v", err)
		return err
	}

	mw := io.MultiWriter(os.Stdout, file)

	handler := slog.NewJSONHandler(mw, &slog.HandlerOptions{Level: LogLevel})
	slog.SetDefault(slog.New(handler))

	return nil
}
