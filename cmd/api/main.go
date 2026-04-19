package main

import (
	"os"

	"haejoong.com-api/internal/api"
	"haejoong.com-api/internal/config"
	"haejoong.com-api/internal/infra"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		os.Exit(1)
	}

	err = infra.InitLogger(cfg.Log)
	if err != nil {
		os.Exit(1)
	}

	appInfra, err := infra.New(cfg)
	if err != nil {
		os.Exit(1)
	}

	err = api.Run(appInfra)
	if err != nil {
		os.Exit(1)
	}
}
