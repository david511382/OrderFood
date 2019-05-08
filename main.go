package main

import (
	"orderfood/src/config"
	"orderfood/src/handler"
	"orderfood/src/logic"
)

var (
	cfg *config.Config
)

// @title Order Food API
// @version 1.0
// @description 訂餐系統

// @BasePath /api/

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	initServer()

	logic.Init(cfg)

	router := handler.Init(isReleaseMode)

	addr := cfg.Domain()

	run(router, addr)
}

func init() {
	err := config.ReadConfig("./src/config/config.yml")
	if err != nil {
		panic(err)
	}

	cfg = config.Get()
}
