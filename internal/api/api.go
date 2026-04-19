package api

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"haejoong.com-api/internal/infra"
)

func Run(appInfra *infra.Infra) error {
	r := gin.Default()

	err := r.Run()

	if err != nil {
		slog.Error("api 서버 가동 실패", "error", err)
	}

	return err
}
